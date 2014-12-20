// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package codedeploy provides a client for AWS CodeDeploy.
package codedeploy

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

// CodeDeploy is a client for AWS CodeDeploy.
type CodeDeploy struct {
	client *aws.JSONClient
}

// New returns a new CodeDeploy client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *CodeDeploy {
	if client == nil {
		client = http.DefaultClient
	}

	service := "codedeploy"
	endpoint, service, region := endpoints.Lookup("codedeploy", region)

	return &CodeDeploy{
		client: &aws.JSONClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			}, Client: client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "CodeDeploy_20141006",
		},
	}
}

// BatchGetApplications is undocumented.
func (c *CodeDeploy) BatchGetApplications(req *BatchGetApplicationsInput) (resp *BatchGetApplicationsOutput, err error) {
	resp = &BatchGetApplicationsOutput{}
	err = c.client.Do("BatchGetApplications", "POST", "/", req, resp)
	return
}

// BatchGetDeployments is undocumented.
func (c *CodeDeploy) BatchGetDeployments(req *BatchGetDeploymentsInput) (resp *BatchGetDeploymentsOutput, err error) {
	resp = &BatchGetDeploymentsOutput{}
	err = c.client.Do("BatchGetDeployments", "POST", "/", req, resp)
	return
}

// CreateApplication is undocumented.
func (c *CodeDeploy) CreateApplication(req *CreateApplicationInput) (resp *CreateApplicationOutput, err error) {
	resp = &CreateApplicationOutput{}
	err = c.client.Do("CreateApplication", "POST", "/", req, resp)
	return
}

// CreateDeployment deploys an application revision to the specified
// deployment group.
func (c *CodeDeploy) CreateDeployment(req *CreateDeploymentInput) (resp *CreateDeploymentOutput, err error) {
	resp = &CreateDeploymentOutput{}
	err = c.client.Do("CreateDeployment", "POST", "/", req, resp)
	return
}

// CreateDeploymentConfig is undocumented.
func (c *CodeDeploy) CreateDeploymentConfig(req *CreateDeploymentConfigInput) (resp *CreateDeploymentConfigOutput, err error) {
	resp = &CreateDeploymentConfigOutput{}
	err = c.client.Do("CreateDeploymentConfig", "POST", "/", req, resp)
	return
}

// CreateDeploymentGroup creates a new deployment group for application
// revisions to be deployed to.
func (c *CodeDeploy) CreateDeploymentGroup(req *CreateDeploymentGroupInput) (resp *CreateDeploymentGroupOutput, err error) {
	resp = &CreateDeploymentGroupOutput{}
	err = c.client.Do("CreateDeploymentGroup", "POST", "/", req, resp)
	return
}

// DeleteApplication is undocumented.
func (c *CodeDeploy) DeleteApplication(req *DeleteApplicationInput) (err error) {
	// NRE
	err = c.client.Do("DeleteApplication", "POST", "/", req, nil)
	return
}

// DeleteDeploymentConfig deletes a deployment configuration. A deployment
// configuration cannot be deleted if it is currently in use. Also,
// predefined configurations cannot be deleted.
func (c *CodeDeploy) DeleteDeploymentConfig(req *DeleteDeploymentConfigInput) (err error) {
	// NRE
	err = c.client.Do("DeleteDeploymentConfig", "POST", "/", req, nil)
	return
}

// DeleteDeploymentGroup is undocumented.
func (c *CodeDeploy) DeleteDeploymentGroup(req *DeleteDeploymentGroupInput) (resp *DeleteDeploymentGroupOutput, err error) {
	resp = &DeleteDeploymentGroupOutput{}
	err = c.client.Do("DeleteDeploymentGroup", "POST", "/", req, resp)
	return
}

// GetApplication is undocumented.
func (c *CodeDeploy) GetApplication(req *GetApplicationInput) (resp *GetApplicationOutput, err error) {
	resp = &GetApplicationOutput{}
	err = c.client.Do("GetApplication", "POST", "/", req, resp)
	return
}

// GetApplicationRevision is undocumented.
func (c *CodeDeploy) GetApplicationRevision(req *GetApplicationRevisionInput) (resp *GetApplicationRevisionOutput, err error) {
	resp = &GetApplicationRevisionOutput{}
	err = c.client.Do("GetApplicationRevision", "POST", "/", req, resp)
	return
}

// GetDeployment is undocumented.
func (c *CodeDeploy) GetDeployment(req *GetDeploymentInput) (resp *GetDeploymentOutput, err error) {
	resp = &GetDeploymentOutput{}
	err = c.client.Do("GetDeployment", "POST", "/", req, resp)
	return
}

// GetDeploymentConfig is undocumented.
func (c *CodeDeploy) GetDeploymentConfig(req *GetDeploymentConfigInput) (resp *GetDeploymentConfigOutput, err error) {
	resp = &GetDeploymentConfigOutput{}
	err = c.client.Do("GetDeploymentConfig", "POST", "/", req, resp)
	return
}

// GetDeploymentGroup is undocumented.
func (c *CodeDeploy) GetDeploymentGroup(req *GetDeploymentGroupInput) (resp *GetDeploymentGroupOutput, err error) {
	resp = &GetDeploymentGroupOutput{}
	err = c.client.Do("GetDeploymentGroup", "POST", "/", req, resp)
	return
}

// GetDeploymentInstance gets information about an Amazon EC2 instance as
// part of a deployment.
func (c *CodeDeploy) GetDeploymentInstance(req *GetDeploymentInstanceInput) (resp *GetDeploymentInstanceOutput, err error) {
	resp = &GetDeploymentInstanceOutput{}
	err = c.client.Do("GetDeploymentInstance", "POST", "/", req, resp)
	return
}

// ListApplicationRevisions lists information about revisions for an
// application.
func (c *CodeDeploy) ListApplicationRevisions(req *ListApplicationRevisionsInput) (resp *ListApplicationRevisionsOutput, err error) {
	resp = &ListApplicationRevisionsOutput{}
	err = c.client.Do("ListApplicationRevisions", "POST", "/", req, resp)
	return
}

// ListApplications lists the applications registered within the AWS user
// account.
func (c *CodeDeploy) ListApplications(req *ListApplicationsInput) (resp *ListApplicationsOutput, err error) {
	resp = &ListApplicationsOutput{}
	err = c.client.Do("ListApplications", "POST", "/", req, resp)
	return
}

// ListDeploymentConfigs lists the deployment configurations within the AWS
// user account.
func (c *CodeDeploy) ListDeploymentConfigs(req *ListDeploymentConfigsInput) (resp *ListDeploymentConfigsOutput, err error) {
	resp = &ListDeploymentConfigsOutput{}
	err = c.client.Do("ListDeploymentConfigs", "POST", "/", req, resp)
	return
}

// ListDeploymentGroups lists the deployment groups for an application
// registered within the AWS user account.
func (c *CodeDeploy) ListDeploymentGroups(req *ListDeploymentGroupsInput) (resp *ListDeploymentGroupsOutput, err error) {
	resp = &ListDeploymentGroupsOutput{}
	err = c.client.Do("ListDeploymentGroups", "POST", "/", req, resp)
	return
}

// ListDeploymentInstances lists the Amazon EC2 instances for a deployment
// within the AWS user account.
func (c *CodeDeploy) ListDeploymentInstances(req *ListDeploymentInstancesInput) (resp *ListDeploymentInstancesOutput, err error) {
	resp = &ListDeploymentInstancesOutput{}
	err = c.client.Do("ListDeploymentInstances", "POST", "/", req, resp)
	return
}

// ListDeployments lists the deployments under a deployment group for an
// application registered within the AWS user account.
func (c *CodeDeploy) ListDeployments(req *ListDeploymentsInput) (resp *ListDeploymentsOutput, err error) {
	resp = &ListDeploymentsOutput{}
	err = c.client.Do("ListDeployments", "POST", "/", req, resp)
	return
}

// RegisterApplicationRevision registers with AWS CodeDeploy a revision for
// the specified application.
func (c *CodeDeploy) RegisterApplicationRevision(req *RegisterApplicationRevisionInput) (err error) {
	// NRE
	err = c.client.Do("RegisterApplicationRevision", "POST", "/", req, nil)
	return
}

// StopDeployment is undocumented.
func (c *CodeDeploy) StopDeployment(req *StopDeploymentInput) (resp *StopDeploymentOutput, err error) {
	resp = &StopDeploymentOutput{}
	err = c.client.Do("StopDeployment", "POST", "/", req, resp)
	return
}

// UpdateApplication is undocumented.
func (c *CodeDeploy) UpdateApplication(req *UpdateApplicationInput) (err error) {
	// NRE
	err = c.client.Do("UpdateApplication", "POST", "/", req, nil)
	return
}

// UpdateDeploymentGroup changes information about an existing deployment
// group.
func (c *CodeDeploy) UpdateDeploymentGroup(req *UpdateDeploymentGroupInput) (resp *UpdateDeploymentGroupOutput, err error) {
	resp = &UpdateDeploymentGroupOutput{}
	err = c.client.Do("UpdateDeploymentGroup", "POST", "/", req, resp)
	return
}

// ApplicationInfo is undocumented.
type ApplicationInfo struct {
	ApplicationID   aws.StringValue   `json:"applicationId,omitempty"`
	ApplicationName aws.StringValue   `json:"applicationName,omitempty"`
	CreateTime      aws.LongTimestamp `json:"createTime,omitempty"`
	LinkedToGitHub  aws.BooleanValue  `json:"linkedToGitHub,omitempty"`
}

// Possible values for CodeDeploy.
const (
	ApplicationRevisionSortByFirstUsedTime = "firstUsedTime"
	ApplicationRevisionSortByLastUsedTime  = "lastUsedTime"
	ApplicationRevisionSortByRegisterTime  = "registerTime"
)

// AutoScalingGroup is undocumented.
type AutoScalingGroup struct {
	Hook aws.StringValue `json:"hook,omitempty"`
	Name aws.StringValue `json:"name,omitempty"`
}

// BatchGetApplicationsInput is undocumented.
type BatchGetApplicationsInput struct {
	ApplicationNames []string `json:"applicationNames,omitempty"`
}

// BatchGetApplicationsOutput is undocumented.
type BatchGetApplicationsOutput struct {
	ApplicationsInfo []ApplicationInfo `json:"applicationsInfo,omitempty"`
}

// BatchGetDeploymentsInput is undocumented.
type BatchGetDeploymentsInput struct {
	DeploymentIDs []string `json:"deploymentIds,omitempty"`
}

// BatchGetDeploymentsOutput is undocumented.
type BatchGetDeploymentsOutput struct {
	DeploymentsInfo []DeploymentInfo `json:"deploymentsInfo,omitempty"`
}

// Possible values for CodeDeploy.
const (
	BundleTypeTAR = "tar"
	BundleTypeTGZ = "tgz"
	BundleTypeZip = "zip"
)

// CreateApplicationInput is undocumented.
type CreateApplicationInput struct {
	ApplicationName aws.StringValue `json:"applicationName"`
}

// CreateApplicationOutput is undocumented.
type CreateApplicationOutput struct {
	ApplicationID aws.StringValue `json:"applicationId,omitempty"`
}

// CreateDeploymentConfigInput is undocumented.
type CreateDeploymentConfigInput struct {
	DeploymentConfigName aws.StringValue      `json:"deploymentConfigName"`
	MinimumHealthyHosts  *MinimumHealthyHosts `json:"minimumHealthyHosts,omitempty"`
}

// CreateDeploymentConfigOutput is undocumented.
type CreateDeploymentConfigOutput struct {
	DeploymentConfigID aws.StringValue `json:"deploymentConfigId,omitempty"`
}

// CreateDeploymentGroupInput is undocumented.
type CreateDeploymentGroupInput struct {
	ApplicationName      aws.StringValue `json:"applicationName"`
	AutoScalingGroups    []string        `json:"autoScalingGroups,omitempty"`
	DeploymentConfigName aws.StringValue `json:"deploymentConfigName,omitempty"`
	DeploymentGroupName  aws.StringValue `json:"deploymentGroupName"`
	EC2TagFilters        []EC2TagFilter  `json:"ec2TagFilters,omitempty"`
	ServiceRoleARN       aws.StringValue `json:"serviceRoleArn,omitempty"`
}

// CreateDeploymentGroupOutput is undocumented.
type CreateDeploymentGroupOutput struct {
	DeploymentGroupID aws.StringValue `json:"deploymentGroupId,omitempty"`
}

// CreateDeploymentInput is undocumented.
type CreateDeploymentInput struct {
	ApplicationName               aws.StringValue   `json:"applicationName"`
	DeploymentConfigName          aws.StringValue   `json:"deploymentConfigName,omitempty"`
	DeploymentGroupName           aws.StringValue   `json:"deploymentGroupName,omitempty"`
	Description                   aws.StringValue   `json:"description,omitempty"`
	IgnoreApplicationStopFailures aws.BooleanValue  `json:"ignoreApplicationStopFailures,omitempty"`
	Revision                      *RevisionLocation `json:"revision,omitempty"`
}

// CreateDeploymentOutput is undocumented.
type CreateDeploymentOutput struct {
	DeploymentID aws.StringValue `json:"deploymentId,omitempty"`
}

// DeleteApplicationInput is undocumented.
type DeleteApplicationInput struct {
	ApplicationName aws.StringValue `json:"applicationName"`
}

// DeleteDeploymentConfigInput is undocumented.
type DeleteDeploymentConfigInput struct {
	DeploymentConfigName aws.StringValue `json:"deploymentConfigName"`
}

// DeleteDeploymentGroupInput is undocumented.
type DeleteDeploymentGroupInput struct {
	ApplicationName     aws.StringValue `json:"applicationName"`
	DeploymentGroupName aws.StringValue `json:"deploymentGroupName"`
}

// DeleteDeploymentGroupOutput is undocumented.
type DeleteDeploymentGroupOutput struct {
	HooksNotCleanedUp []AutoScalingGroup `json:"hooksNotCleanedUp,omitempty"`
}

// DeploymentConfigInfo is undocumented.
type DeploymentConfigInfo struct {
	CreateTime           aws.LongTimestamp    `json:"createTime,omitempty"`
	DeploymentConfigID   aws.StringValue      `json:"deploymentConfigId,omitempty"`
	DeploymentConfigName aws.StringValue      `json:"deploymentConfigName,omitempty"`
	MinimumHealthyHosts  *MinimumHealthyHosts `json:"minimumHealthyHosts,omitempty"`
}

// Possible values for CodeDeploy.
const (
	DeploymentCreatorAutoscaling = "autoscaling"
	DeploymentCreatorUser        = "user"
)

// DeploymentGroupInfo is undocumented.
type DeploymentGroupInfo struct {
	ApplicationName      aws.StringValue    `json:"applicationName,omitempty"`
	AutoScalingGroups    []AutoScalingGroup `json:"autoScalingGroups,omitempty"`
	DeploymentConfigName aws.StringValue    `json:"deploymentConfigName,omitempty"`
	DeploymentGroupID    aws.StringValue    `json:"deploymentGroupId,omitempty"`
	DeploymentGroupName  aws.StringValue    `json:"deploymentGroupName,omitempty"`
	EC2TagFilters        []EC2TagFilter     `json:"ec2TagFilters,omitempty"`
	ServiceRoleARN       aws.StringValue    `json:"serviceRoleArn,omitempty"`
	TargetRevision       *RevisionLocation  `json:"targetRevision,omitempty"`
}

// DeploymentInfo is undocumented.
type DeploymentInfo struct {
	ApplicationName               aws.StringValue     `json:"applicationName,omitempty"`
	CompleteTime                  aws.LongTimestamp   `json:"completeTime,omitempty"`
	CreateTime                    aws.LongTimestamp   `json:"createTime,omitempty"`
	Creator                       aws.StringValue     `json:"creator,omitempty"`
	DeploymentConfigName          aws.StringValue     `json:"deploymentConfigName,omitempty"`
	DeploymentGroupName           aws.StringValue     `json:"deploymentGroupName,omitempty"`
	DeploymentID                  aws.StringValue     `json:"deploymentId,omitempty"`
	DeploymentOverview            *DeploymentOverview `json:"deploymentOverview,omitempty"`
	Description                   aws.StringValue     `json:"description,omitempty"`
	ErrorInformation              *ErrorInformation   `json:"errorInformation,omitempty"`
	IgnoreApplicationStopFailures aws.BooleanValue    `json:"ignoreApplicationStopFailures,omitempty"`
	Revision                      *RevisionLocation   `json:"revision,omitempty"`
	StartTime                     aws.LongTimestamp   `json:"startTime,omitempty"`
	Status                        aws.StringValue     `json:"status,omitempty"`
}

// DeploymentOverview is undocumented.
type DeploymentOverview struct {
	Failed     aws.LongValue `json:"Failed,omitempty"`
	InProgress aws.LongValue `json:"InProgress,omitempty"`
	Pending    aws.LongValue `json:"Pending,omitempty"`
	Skipped    aws.LongValue `json:"Skipped,omitempty"`
	Succeeded  aws.LongValue `json:"Succeeded,omitempty"`
}

// Possible values for CodeDeploy.
const (
	DeploymentStatusCreated    = "Created"
	DeploymentStatusFailed     = "Failed"
	DeploymentStatusInProgress = "InProgress"
	DeploymentStatusQueued     = "Queued"
	DeploymentStatusStopped    = "Stopped"
	DeploymentStatusSucceeded  = "Succeeded"
)

// Diagnostics is undocumented.
type Diagnostics struct {
	ErrorCode  aws.StringValue `json:"errorCode,omitempty"`
	LogTail    aws.StringValue `json:"logTail,omitempty"`
	Message    aws.StringValue `json:"message,omitempty"`
	ScriptName aws.StringValue `json:"scriptName,omitempty"`
}

// EC2TagFilter is undocumented.
type EC2TagFilter struct {
	Key   aws.StringValue `json:"Key,omitempty"`
	Type  aws.StringValue `json:"Type,omitempty"`
	Value aws.StringValue `json:"Value,omitempty"`
}

// Possible values for CodeDeploy.
const (
	EC2TagFilterTypeKeyAndValue = "KEY_AND_VALUE"
	EC2TagFilterTypeKeyOnly     = "KEY_ONLY"
	EC2TagFilterTypeValueOnly   = "VALUE_ONLY"
)

// Possible values for CodeDeploy.
const (
	ErrorCodeApplicationMissing       = "APPLICATION_MISSING"
	ErrorCodeDeploymentGroupMissing   = "DEPLOYMENT_GROUP_MISSING"
	ErrorCodeHealthConstraints        = "HEALTH_CONSTRAINTS"
	ErrorCodeHealthConstraintsInvalid = "HEALTH_CONSTRAINTS_INVALID"
	ErrorCodeIAMRoleMissing           = "IAM_ROLE_MISSING"
	ErrorCodeIAMRolePermissions       = "IAM_ROLE_PERMISSIONS"
	ErrorCodeInternalError            = "INTERNAL_ERROR"
	ErrorCodeNoInstances              = "NO_INSTANCES"
	ErrorCodeOverMaxInstances         = "OVER_MAX_INSTANCES"
	ErrorCodeRevisionMissing          = "REVISION_MISSING"
	ErrorCodeTimeout                  = "TIMEOUT"
)

// ErrorInformation is undocumented.
type ErrorInformation struct {
	Code    aws.StringValue `json:"code,omitempty"`
	Message aws.StringValue `json:"message,omitempty"`
}

// GenericRevisionInfo is undocumented.
type GenericRevisionInfo struct {
	DeploymentGroups []string          `json:"deploymentGroups,omitempty"`
	Description      aws.StringValue   `json:"description,omitempty"`
	FirstUsedTime    aws.LongTimestamp `json:"firstUsedTime,omitempty"`
	LastUsedTime     aws.LongTimestamp `json:"lastUsedTime,omitempty"`
	RegisterTime     aws.LongTimestamp `json:"registerTime,omitempty"`
}

// GetApplicationInput is undocumented.
type GetApplicationInput struct {
	ApplicationName aws.StringValue `json:"applicationName"`
}

// GetApplicationOutput is undocumented.
type GetApplicationOutput struct {
	Application *ApplicationInfo `json:"application,omitempty"`
}

// GetApplicationRevisionInput is undocumented.
type GetApplicationRevisionInput struct {
	ApplicationName aws.StringValue   `json:"applicationName"`
	Revision        *RevisionLocation `json:"revision"`
}

// GetApplicationRevisionOutput is undocumented.
type GetApplicationRevisionOutput struct {
	ApplicationName aws.StringValue      `json:"applicationName,omitempty"`
	Revision        *RevisionLocation    `json:"revision,omitempty"`
	RevisionInfo    *GenericRevisionInfo `json:"revisionInfo,omitempty"`
}

// GetDeploymentConfigInput is undocumented.
type GetDeploymentConfigInput struct {
	DeploymentConfigName aws.StringValue `json:"deploymentConfigName"`
}

// GetDeploymentConfigOutput is undocumented.
type GetDeploymentConfigOutput struct {
	DeploymentConfigInfo *DeploymentConfigInfo `json:"deploymentConfigInfo,omitempty"`
}

// GetDeploymentGroupInput is undocumented.
type GetDeploymentGroupInput struct {
	ApplicationName     aws.StringValue `json:"applicationName"`
	DeploymentGroupName aws.StringValue `json:"deploymentGroupName"`
}

// GetDeploymentGroupOutput is undocumented.
type GetDeploymentGroupOutput struct {
	DeploymentGroupInfo *DeploymentGroupInfo `json:"deploymentGroupInfo,omitempty"`
}

// GetDeploymentInput is undocumented.
type GetDeploymentInput struct {
	DeploymentID aws.StringValue `json:"deploymentId"`
}

// GetDeploymentInstanceInput is undocumented.
type GetDeploymentInstanceInput struct {
	DeploymentID aws.StringValue `json:"deploymentId"`
	InstanceID   aws.StringValue `json:"instanceId"`
}

// GetDeploymentInstanceOutput is undocumented.
type GetDeploymentInstanceOutput struct {
	InstanceSummary *InstanceSummary `json:"instanceSummary,omitempty"`
}

// GetDeploymentOutput is undocumented.
type GetDeploymentOutput struct {
	DeploymentInfo *DeploymentInfo `json:"deploymentInfo,omitempty"`
}

// GitHubLocation is undocumented.
type GitHubLocation struct {
	CommitID   aws.StringValue `json:"commitId,omitempty"`
	Repository aws.StringValue `json:"repository,omitempty"`
}

// Possible values for CodeDeploy.
const (
	InstanceStatusFailed     = "Failed"
	InstanceStatusInProgress = "InProgress"
	InstanceStatusPending    = "Pending"
	InstanceStatusSkipped    = "Skipped"
	InstanceStatusSucceeded  = "Succeeded"
	InstanceStatusUnknown    = "Unknown"
)

// InstanceSummary is undocumented.
type InstanceSummary struct {
	DeploymentID    aws.StringValue   `json:"deploymentId,omitempty"`
	InstanceID      aws.StringValue   `json:"instanceId,omitempty"`
	LastUpdatedAt   aws.LongTimestamp `json:"lastUpdatedAt,omitempty"`
	LifecycleEvents []LifecycleEvent  `json:"lifecycleEvents,omitempty"`
	Status          aws.StringValue   `json:"status,omitempty"`
}

// Possible values for CodeDeploy.
const (
	LifecycleErrorCodeScriptFailed        = "ScriptFailed"
	LifecycleErrorCodeScriptMissing       = "ScriptMissing"
	LifecycleErrorCodeScriptNotExecutable = "ScriptNotExecutable"
	LifecycleErrorCodeScriptTimedOut      = "ScriptTimedOut"
	LifecycleErrorCodeSuccess             = "Success"
	LifecycleErrorCodeUnknownError        = "UnknownError"
)

// LifecycleEvent is undocumented.
type LifecycleEvent struct {
	Diagnostics        *Diagnostics      `json:"diagnostics,omitempty"`
	EndTime            aws.LongTimestamp `json:"endTime,omitempty"`
	LifecycleEventName aws.StringValue   `json:"lifecycleEventName,omitempty"`
	StartTime          aws.LongTimestamp `json:"startTime,omitempty"`
	Status             aws.StringValue   `json:"status,omitempty"`
}

// Possible values for CodeDeploy.
const (
	LifecycleEventStatusFailed     = "Failed"
	LifecycleEventStatusInProgress = "InProgress"
	LifecycleEventStatusPending    = "Pending"
	LifecycleEventStatusSkipped    = "Skipped"
	LifecycleEventStatusSucceeded  = "Succeeded"
	LifecycleEventStatusUnknown    = "Unknown"
)

// ListApplicationRevisionsInput is undocumented.
type ListApplicationRevisionsInput struct {
	ApplicationName aws.StringValue `json:"applicationName"`
	Deployed        aws.StringValue `json:"deployed,omitempty"`
	NextToken       aws.StringValue `json:"nextToken,omitempty"`
	S3Bucket        aws.StringValue `json:"s3Bucket,omitempty"`
	S3KeyPrefix     aws.StringValue `json:"s3KeyPrefix,omitempty"`
	SortBy          aws.StringValue `json:"sortBy,omitempty"`
	SortOrder       aws.StringValue `json:"sortOrder,omitempty"`
}

// ListApplicationRevisionsOutput is undocumented.
type ListApplicationRevisionsOutput struct {
	NextToken aws.StringValue    `json:"nextToken,omitempty"`
	Revisions []RevisionLocation `json:"revisions,omitempty"`
}

// ListApplicationsInput is undocumented.
type ListApplicationsInput struct {
	NextToken aws.StringValue `json:"nextToken,omitempty"`
}

// ListApplicationsOutput is undocumented.
type ListApplicationsOutput struct {
	Applications []string        `json:"applications,omitempty"`
	NextToken    aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentConfigsInput is undocumented.
type ListDeploymentConfigsInput struct {
	NextToken aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentConfigsOutput is undocumented.
type ListDeploymentConfigsOutput struct {
	DeploymentConfigsList []string        `json:"deploymentConfigsList,omitempty"`
	NextToken             aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentGroupsInput is undocumented.
type ListDeploymentGroupsInput struct {
	ApplicationName aws.StringValue `json:"applicationName"`
	NextToken       aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentGroupsOutput is undocumented.
type ListDeploymentGroupsOutput struct {
	ApplicationName  aws.StringValue `json:"applicationName,omitempty"`
	DeploymentGroups []string        `json:"deploymentGroups,omitempty"`
	NextToken        aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentInstancesInput is undocumented.
type ListDeploymentInstancesInput struct {
	DeploymentID         aws.StringValue `json:"deploymentId"`
	InstanceStatusFilter []string        `json:"instanceStatusFilter,omitempty"`
	NextToken            aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentInstancesOutput is undocumented.
type ListDeploymentInstancesOutput struct {
	InstancesList []string        `json:"instancesList,omitempty"`
	NextToken     aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentsInput is undocumented.
type ListDeploymentsInput struct {
	ApplicationName     aws.StringValue `json:"applicationName,omitempty"`
	CreateTimeRange     *TimeRange      `json:"createTimeRange,omitempty"`
	DeploymentGroupName aws.StringValue `json:"deploymentGroupName,omitempty"`
	IncludeOnlyStatuses []string        `json:"includeOnlyStatuses,omitempty"`
	NextToken           aws.StringValue `json:"nextToken,omitempty"`
}

// ListDeploymentsOutput is undocumented.
type ListDeploymentsOutput struct {
	Deployments []string        `json:"deployments,omitempty"`
	NextToken   aws.StringValue `json:"nextToken,omitempty"`
}

// Possible values for CodeDeploy.
const (
	ListStateFilterActionExclude = "exclude"
	ListStateFilterActionIgnore  = "ignore"
	ListStateFilterActionInclude = "include"
)

// MinimumHealthyHosts is undocumented.
type MinimumHealthyHosts struct {
	Type  aws.StringValue  `json:"type,omitempty"`
	Value aws.IntegerValue `json:"value,omitempty"`
}

// Possible values for CodeDeploy.
const (
	MinimumHealthyHostsTypeFleetPercent = "FLEET_PERCENT"
	MinimumHealthyHostsTypeHostCount    = "HOST_COUNT"
)

// RegisterApplicationRevisionInput is undocumented.
type RegisterApplicationRevisionInput struct {
	ApplicationName aws.StringValue   `json:"applicationName"`
	Description     aws.StringValue   `json:"description,omitempty"`
	Revision        *RevisionLocation `json:"revision"`
}

// RevisionLocation is undocumented.
type RevisionLocation struct {
	GitHubLocation *GitHubLocation `json:"gitHubLocation,omitempty"`
	RevisionType   aws.StringValue `json:"revisionType,omitempty"`
	S3Location     *S3Location     `json:"s3Location,omitempty"`
}

// Possible values for CodeDeploy.
const (
	RevisionLocationTypeGitHub = "GitHub"
	RevisionLocationTypeS3     = "S3"
)

// S3Location is undocumented.
type S3Location struct {
	Bucket     aws.StringValue `json:"bucket,omitempty"`
	BundleType aws.StringValue `json:"bundleType,omitempty"`
	ETag       aws.StringValue `json:"eTag,omitempty"`
	Key        aws.StringValue `json:"key,omitempty"`
	Version    aws.StringValue `json:"version,omitempty"`
}

// Possible values for CodeDeploy.
const (
	SortOrderAscending  = "ascending"
	SortOrderDescending = "descending"
)

// StopDeploymentInput is undocumented.
type StopDeploymentInput struct {
	DeploymentID aws.StringValue `json:"deploymentId"`
}

// StopDeploymentOutput is undocumented.
type StopDeploymentOutput struct {
	Status        aws.StringValue `json:"status,omitempty"`
	StatusMessage aws.StringValue `json:"statusMessage,omitempty"`
}

// Possible values for CodeDeploy.
const (
	StopStatusPending   = "Pending"
	StopStatusSucceeded = "Succeeded"
)

// TimeRange is undocumented.
type TimeRange struct {
	End   aws.LongTimestamp `json:"end,omitempty"`
	Start aws.LongTimestamp `json:"start,omitempty"`
}

// UpdateApplicationInput is undocumented.
type UpdateApplicationInput struct {
	ApplicationName    aws.StringValue `json:"applicationName,omitempty"`
	NewApplicationName aws.StringValue `json:"newApplicationName,omitempty"`
}

// UpdateDeploymentGroupInput is undocumented.
type UpdateDeploymentGroupInput struct {
	ApplicationName            aws.StringValue `json:"applicationName"`
	AutoScalingGroups          []string        `json:"autoScalingGroups,omitempty"`
	CurrentDeploymentGroupName aws.StringValue `json:"currentDeploymentGroupName"`
	DeploymentConfigName       aws.StringValue `json:"deploymentConfigName,omitempty"`
	EC2TagFilters              []EC2TagFilter  `json:"ec2TagFilters,omitempty"`
	NewDeploymentGroupName     aws.StringValue `json:"newDeploymentGroupName,omitempty"`
	ServiceRoleARN             aws.StringValue `json:"serviceRoleArn,omitempty"`
}

// UpdateDeploymentGroupOutput is undocumented.
type UpdateDeploymentGroupOutput struct {
	HooksNotCleanedUp []AutoScalingGroup `json:"hooksNotCleanedUp,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
