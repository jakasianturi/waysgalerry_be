package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadFile(next http.HandlerFunc, formInputName string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upload file
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, fileHeader, err := r.FormFile(formInputName)

		if err != nil {
			// fmt.Println(err)
			// add filename to ctx
			ctx := context.WithValue(r.Context(), "Error", true)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		defer file.Close()
		// cloudinary config
		var ctxcld = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")
		fileInput, _ := fileHeader.Open()
		defer file.Close()
		// looping data files for sent one by one to cloudinary
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctxcld, fileInput, uploader.UploadParams{Folder: "b42-waysgallery"})
		if err != nil {
			fmt.Println(err.Error())
		}
		data := resp.SecureURL
		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)
		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFile", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
