// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ImportImageInput struct {
	_ struct{} `type:"structure"`

	// The architecture of the virtual machine.
	//
	// Valid values: i386 | x86_64 | arm64
	Architecture *string `type:"string"`

	// The client-specific data.
	ClientData *Data `type:"structure"`

	// The token to enable idempotency for VM import requests.
	ClientToken *string `type:"string"`

	// A description string for the import image task.
	Description *string `type:"string"`

	// Information about the disk containers.
	DiskContainers []ImageDiskContainer `locationName:"DiskContainer" locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// Specifies whether the destination AMI of the imported image should be encrypted.
	// The default CMK for EBS is used unless you specify a non-default AWS Key
	// Management Service (AWS KMS) CMK using KmsKeyId. For more information, see
	// Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool `type:"boolean"`

	// The target hypervisor platform.
	//
	// Valid values: xen
	Hypervisor *string `type:"string"`

	// An identifier for the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) to use when creating the encrypted AMI. This parameter is
	// only required if you want to use a non-default CMK; if this parameter is
	// not specified, the default CMK for EBS is used. If a KmsKeyId is specified,
	// the Encrypted flag must also be set.
	//
	// The CMK identifier may be provided in any of the following formats:
	//
	//    * Key ID
	//
	//    * Key alias. The alias ARN contains the arn:aws:kms namespace, followed
	//    by the Region of the CMK, the AWS account ID of the CMK owner, the alias
	//    namespace, and then the CMK alias. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	//    * ARN using key ID. The ID ARN contains the arn:aws:kms namespace, followed
	//    by the Region of the CMK, the AWS account ID of the CMK owner, the key
	//    namespace, and then the CMK ID. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//    * ARN using key alias. The alias ARN contains the arn:aws:kms namespace,
	//    followed by the Region of the CMK, the AWS account ID of the CMK owner,
	//    the alias namespace, and then the CMK alias. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS parses KmsKeyId asynchronously, meaning that the action you call may
	// appear to complete even though you provided an invalid identifier. This action
	// will eventually report failure.
	//
	// The specified CMK must exist in the Region that the AMI is being copied to.
	//
	// Amazon EBS does not support asymmetric CMKs.
	KmsKeyId *string `type:"string"`

	// The ARNs of the license configurations.
	LicenseSpecifications []ImportImageLicenseConfigurationRequest `locationNameList:"item" type:"list"`

	// The license type to be used for the Amazon Machine Image (AMI) after importing.
	//
	// By default, we detect the source-system operating system (OS) and apply the
	// appropriate license. Specify AWS to replace the source-system license with
	// an AWS license, if appropriate. Specify BYOL to retain the source-system
	// license, if appropriate.
	//
	// To use BYOL, you must have existing licenses with rights to use these licenses
	// in a third party cloud, such as AWS. For more information, see Prerequisites
	// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html#prerequisites-image)
	// in the VM Import/Export User Guide.
	LicenseType *string `type:"string"`

	// The operating system of the virtual machine.
	//
	// Valid values: Windows | Linux
	Platform *string `type:"string"`

	// The name of the role to use when not using the default role, 'vmimport'.
	RoleName *string `type:"string"`
}

// String returns the string representation
func (s ImportImageInput) String() string {
	return awsutil.Prettify(s)
}

type ImportImageOutput struct {
	_ struct{} `type:"structure"`

	// The architecture of the virtual machine.
	Architecture *string `locationName:"architecture" type:"string"`

	// A description of the import task.
	Description *string `locationName:"description" type:"string"`

	// Indicates whether the AMI is encypted.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// The target hypervisor of the import task.
	Hypervisor *string `locationName:"hypervisor" type:"string"`

	// The ID of the Amazon Machine Image (AMI) created by the import task.
	ImageId *string `locationName:"imageId" type:"string"`

	// The task ID of the import image task.
	ImportTaskId *string `locationName:"importTaskId" type:"string"`

	// The identifier for the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) that was used to create the encrypted AMI.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The ARNs of the license configurations.
	LicenseSpecifications []ImportImageLicenseConfigurationResponse `locationName:"licenseSpecifications" locationNameList:"item" type:"list"`

	// The license type of the virtual machine.
	LicenseType *string `locationName:"licenseType" type:"string"`

	// The operating system of the virtual machine.
	Platform *string `locationName:"platform" type:"string"`

	// The progress of the task.
	Progress *string `locationName:"progress" type:"string"`

	// Information about the snapshots.
	SnapshotDetails []SnapshotDetail `locationName:"snapshotDetailSet" locationNameList:"item" type:"list"`

	// A brief status of the task.
	Status *string `locationName:"status" type:"string"`

	// A detailed status message of the import task.
	StatusMessage *string `locationName:"statusMessage" type:"string"`
}

// String returns the string representation
func (s ImportImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportImage = "ImportImage"

// ImportImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Import single or multi-volume disk images or EBS snapshots into an Amazon
// Machine Image (AMI). For more information, see Importing a VM as an Image
// Using VM Import/Export (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html)
// in the VM Import/Export User Guide.
//
//    // Example sending a request using ImportImageRequest.
//    req := client.ImportImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportImage
func (c *Client) ImportImageRequest(input *ImportImageInput) ImportImageRequest {
	op := &aws.Operation{
		Name:       opImportImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportImageInput{}
	}

	req := c.newRequest(op, input, &ImportImageOutput{})
	return ImportImageRequest{Request: req, Input: input, Copy: c.ImportImageRequest}
}

// ImportImageRequest is the request type for the
// ImportImage API operation.
type ImportImageRequest struct {
	*aws.Request
	Input *ImportImageInput
	Copy  func(*ImportImageInput) ImportImageRequest
}

// Send marshals and sends the ImportImage API request.
func (r ImportImageRequest) Send(ctx context.Context) (*ImportImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportImageResponse{
		ImportImageOutput: r.Request.Data.(*ImportImageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportImageResponse is the response type for the
// ImportImage API operation.
type ImportImageResponse struct {
	*ImportImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportImage request.
func (r *ImportImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
