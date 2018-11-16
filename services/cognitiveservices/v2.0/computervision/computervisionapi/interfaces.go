package computervisionapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	AnalyzeImage(ctx context.Context, imageURL computervision.ImageURL, visualFeatures []computervision.VisualFeatureTypes, details []computervision.Details, language string) (result computervision.ImageAnalysis, err error)
	AnalyzeImageByDomain(ctx context.Context, model string, imageURL computervision.ImageURL, language string) (result computervision.DomainModelResults, err error)
	AnalyzeImageByDomainInStream(ctx context.Context, model string, imageParameter io.ReadCloser, language string) (result computervision.DomainModelResults, err error)
	AnalyzeImageInStream(ctx context.Context, imageParameter io.ReadCloser, visualFeatures []computervision.VisualFeatureTypes, details []computervision.Details, language string) (result computervision.ImageAnalysis, err error)
	DescribeImage(ctx context.Context, imageURL computervision.ImageURL, maxCandidates *int32, language string) (result computervision.ImageDescription, err error)
	DescribeImageInStream(ctx context.Context, imageParameter io.ReadCloser, maxCandidates *int32, language string) (result computervision.ImageDescription, err error)
	GenerateThumbnail(ctx context.Context, width int32, height int32, imageURL computervision.ImageURL, smartCropping *bool) (result computervision.ReadCloser, err error)
	GenerateThumbnailInStream(ctx context.Context, width int32, height int32, imageParameter io.ReadCloser, smartCropping *bool) (result computervision.ReadCloser, err error)
	GetTextOperationResult(ctx context.Context, operationID string) (result computervision.TextOperationResult, err error)
	ListModels(ctx context.Context) (result computervision.ListModelsResult, err error)
	RecognizePrintedText(ctx context.Context, detectOrientation bool, imageURL computervision.ImageURL, language computervision.OcrLanguages) (result computervision.OcrResult, err error)
	RecognizePrintedTextInStream(ctx context.Context, detectOrientation bool, imageParameter io.ReadCloser, language computervision.OcrLanguages) (result computervision.OcrResult, err error)
	RecognizeText(ctx context.Context, imageURL computervision.ImageURL, mode computervision.TextRecognitionMode) (result autorest.Response, err error)
	RecognizeTextInStream(ctx context.Context, imageParameter io.ReadCloser, mode computervision.TextRecognitionMode) (result autorest.Response, err error)
	TagImage(ctx context.Context, imageURL computervision.ImageURL, language string) (result computervision.TagResult, err error)
	TagImageInStream(ctx context.Context, imageParameter io.ReadCloser, language string) (result computervision.TagResult, err error)
}

var _ BaseClientAPI = (*computervision.BaseClient)(nil)