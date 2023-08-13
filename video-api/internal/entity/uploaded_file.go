package entity

type UploadedFile struct {
	bucket   string
	filename string
}

func NewUploadedFile(
	bucket string,
	filename string,
) UploadedFile {
	return UploadedFile{
		bucket:   bucket,
		filename: filename,
	}
}

func (uploadedfile *UploadedFile) GetBucket() string {
	return uploadedfile.bucket
}

func (uploadedfile *UploadedFile) GetFilename() string {
	return uploadedfile.filename
}

func (uploadedfile *UploadedFile) SetBucket(bucket string) *UploadedFile {
	uploadedfile.bucket = bucket
	return uploadedfile
}

func (uploadedfile *UploadedFile) SetFilename(filename string) *UploadedFile {
	uploadedfile.filename = filename
	return uploadedfile
}
