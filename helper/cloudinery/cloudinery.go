package cloudinery

import (
	"context"

	cloud "github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadToCloudinary(cld *cloud.Cloudinary, imagePath string) (string, error) {
	uploadParams := uploader.UploadParams{
		PublicID: "desired_public_id",
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), imagePath, uploadParams)
	if err != nil {
		return "", err
	}
	return uploadResult.URL, nil
}
