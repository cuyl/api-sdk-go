// Copyright (c) 2020 Smartling
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package smartling

import (
	"fmt"
)

const (
	endpointUploadFile = "/files-api/v2/projects/%s/file"
)

type FileUploadResult struct {
	Overwritten bool
	StringCount int
	WordCount   int
}

// DownloadFile downloads original file from project.
func (client *Client) UploadFile(
	projectID string,
	request FileUploadRequest,
) (*FileUploadResult, error) {
	var result FileUploadResult

	form, err := request.GetForm()
	if err != nil {
		return nil, fmt.Errorf("failed to create file upload form: %w", err)
	}

	err = form.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close upload file form: %w", err)
	}

	_, _, err = client.Post(
		fmt.Sprintf(endpointUploadFile, projectID),
		form.Bytes(),
		&result,
		ContentTypeOption(form.GetContentType()),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to upload original file: %w",
			err,
		)
	}

	return &result, nil
}
