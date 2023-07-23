package entity

type UploadedInfo struct {
	Filename string
	Bucket   string
}

func NewUploadedInfo(filename string, bucket string) *UploadedInfo {
	return &UploadedInfo{
		Filename: filename,
		Bucket:   bucket,
	}
}
