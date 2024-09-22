package uploader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

// streamFileUpload uploads an object via a stream.
func StreamFileUpload(w io.Writer, bucket, object string, file []byte) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	buf := bytes.NewBuffer(file)

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err = io.Copy(wc, buf); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	fmt.Fprintf(w, "%v uploaded to %v.\n", object, bucket)

	return nil
}

type UploadFileResponse struct {
	Generation int64
	ObjectName string
	Err        error
}

// uploadFile uploads an object.
func UploadFile(w io.Writer, bucket, object string) UploadFileResponse {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return UploadFileResponse{0, "", fmt.Errorf("storage.NewClient: %w", err)}
	}
	defer client.Close()

	// Open local file.
	filePath := fmt.Sprintf("uploads/%v", object)

	f, err := os.Open(filePath)
	if err != nil {
		return UploadFileResponse{0, "", fmt.Errorf("os.Open: %w", err)}
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(bucket).Object(object)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	// For an object that does not yet exist, set the DoesNotExist precondition.
	// o = o.If(storage.Conditions{DoesNotExist: false})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return UploadFileResponse{0, "", fmt.Errorf("object.Attrs: %w", err)}
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return UploadFileResponse{0, "", fmt.Errorf("io.Copy: %w", err)}
	}
	if err := wc.Close(); err != nil {
		return UploadFileResponse{0, "", fmt.Errorf("Writer.Close: %w", err)}
	}
	fmt.Fprintf(w, "Blob %v uploaded.\n", object)

	generationId, objectName := attrs.Generation, attrs.Name

	return UploadFileResponse{Generation: generationId, ObjectName: objectName, Err: nil}
}
