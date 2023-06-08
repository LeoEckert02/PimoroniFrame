package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/adrium/goheif"
	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// Skip Writer for exif writing
type writerSkipper struct {
	w           io.Writer
	bytesToSkip int
}

func (w *writerSkipper) Write(data []byte) (int, error) {
	if w.bytesToSkip <= 0 {
		return w.w.Write(data)
	}

	if dataLen := len(data); dataLen < w.bytesToSkip {
		w.bytesToSkip -= dataLen
		return dataLen, nil
	}

	if n, err := w.w.Write(data[w.bytesToSkip:]); err == nil {
		n += w.bytesToSkip
		w.bytesToSkip = 0
		return n, nil
	} else {
		return n, err
	}
}

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPut, // Change the method to POST
			Path:   "/api/convert", // Change the path to /api/upload
			Handler: func(c echo.Context) error {
				// Get the uploaded file from the request
				file, err := c.FormFile("image")
				if err != nil {
					return c.JSON(http.StatusBadRequest, "Failed to get the image file")
				}

				// Open the uploaded file
				src, err := file.Open()
				if err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to open the image file")
				}
				defer src.Close()

				// Extract exif to add back in after conversion
				exif, err := goheif.ExtractExif(src)
				if err != nil {
					return err
				}

				// Create a buffer to store the modified image data
				modifiedImageBuf := new(bytes.Buffer)

				img, err := goheif.Decode(src)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to decode the image")
				}

				w, _ := newWriterExif(modifiedImageBuf, exif)
				if err := imaging.Encode(w, img, imaging.JPEG, imaging.JPEGQuality(20)); err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to encode the modified image")
				}

				// Set the response headers to indicate that the response contains an image
				c.Response().Header().Set("Content-Type", "image/jpeg")

				// Write the modified image data from the buffer to the response
				if _, err := modifiedImageBuf.WriteTo(c.Response().Writer); err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to write the modified image to the response")
				}

				return nil
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPut,  // Change the method to POST
			Path:   "/api/compress", // Change the path to /api/upload
			Handler: func(c echo.Context) error {
				// Get the uploaded file from the request
				file, err := c.FormFile("image")
				if err != nil {
					return c.JSON(http.StatusBadRequest, "Failed to get the image file")
				}

				// Open the uploaded file
				src, err := file.Open()
				if err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to open the image file")
				}
				defer src.Close()

				// Decode the HEIC image
				img, err := imaging.Decode(src)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to decode the image")
				}

				// Create a buffer to store the compressed image data
				compressedImgBuf := new(bytes.Buffer)

				// Compress the HEIC image to JPEG format
				err = imaging.Encode(compressedImgBuf, img, imaging.JPEG, imaging.JPEGQuality(20))
				if err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to compress the image")
				}

				// Set the response headers to indicate that the response contains an image
				c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")

				// Write the compressed image data from the buffer to the response
				_, err = io.Copy(c.Response().Writer, compressedImgBuf)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, "Failed to write the image to the response")
				}

				return nil
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})

	/*app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		// overwrite the newly submitted "posts" record status to pending
		if e.Record.Collection().Name == "images" {
			image := e.Record.Get("image")
			imageType := reflect.TypeOf(image).String()
			log.Println(e.Collection())
		}
		return nil
	})
	*/
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func newWriterExif(w io.Writer, exif []byte) (io.Writer, error) {
	writer := &writerSkipper{w, 2}
	soi := []byte{0xff, 0xd8}
	if _, err := w.Write(soi); err != nil {
		return nil, err
	}

	if exif != nil {
		app1Marker := 0xe1
		markerlen := 2 + len(exif)
		marker := []byte{0xff, uint8(app1Marker), uint8(markerlen >> 8), uint8(markerlen & 0xff)}
		if _, err := w.Write(marker); err != nil {
			return nil, err
		}

		if _, err := w.Write(exif); err != nil {
			return nil, err
		}
	}

	return writer, nil
}
