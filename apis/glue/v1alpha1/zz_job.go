/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// JobParameters defines the desired state of Job
type JobParameters struct {
	// Region is which region the Job will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// This parameter is deprecated. Use MaxCapacity instead.
	//
	// The number of AWS Glue data processing units (DPUs) to allocate to this Job.
	// You can allocate from 2 to 100 DPUs; the default is 10. A DPU is a relative
	// measure of processing power that consists of 4 vCPUs of compute capacity
	// and 16 GB of memory. For more information, see the AWS Glue pricing page
	// (https://aws.amazon.com/glue/pricing/).
	AllocatedCapacity *int64 `json:"allocatedCapacity,omitempty"`
	// The JobCommand that executes this job.
	// +kubebuilder:validation:Required
	Command *JobCommand `json:"command"`
	// The connections used for this job.
	Connections *ConnectionsList `json:"connections,omitempty"`
	// The default arguments for this job.
	//
	// You can specify arguments here that your own job-execution script consumes,
	// as well as arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own Job arguments,
	// see the Calling AWS Glue APIs in Python (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html)
	// topic in the developer guide.
	//
	// For information about the key-value pairs that AWS Glue consumes to set up
	// your job, see the Special Parameters Used by AWS Glue (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	// topic in the developer guide.
	DefaultArguments map[string]*string `json:"defaultArguments,omitempty"`
	// Description of the job being defined.
	Description *string `json:"description,omitempty"`
	// An ExecutionProperty specifying the maximum number of concurrent runs allowed
	// for this job.
	ExecutionProperty *ExecutionProperty `json:"executionProperty,omitempty"`
	// Glue version determines the versions of Apache Spark and Python that AWS
	// Glue supports. The Python version indicates the version supported for jobs
	// of type Spark.
	//
	// For more information about the available AWS Glue versions and corresponding
	// Spark and Python versions, see Glue version (https://docs.aws.amazon.com/glue/latest/dg/add-job.html)
	// in the developer guide.
	//
	// Jobs that are created without specifying a Glue version default to Glue 0.9.
	GlueVersion *string `json:"glueVersion,omitempty"`
	// This field is reserved for future use.
	LogURI *string `json:"logURI,omitempty"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated
	// when this job runs. A DPU is a relative measure of processing power that
	// consists of 4 vCPUs of compute capacity and 16 GB of memory. For more information,
	// see the AWS Glue pricing page (https://aws.amazon.com/glue/pricing/).
	//
	// Do not set Max Capacity if using WorkerType and NumberOfWorkers.
	//
	// The value that can be allocated for MaxCapacity depends on whether you are
	// running a Python shell job or an Apache Spark ETL job:
	//
	//    * When you specify a Python shell job (JobCommand.Name="pythonshell"),
	//    you can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	//
	//    * When you specify an Apache Spark ETL job (JobCommand.Name="glueetl")
	//    or Apache Spark streaming ETL job (JobCommand.Name="gluestreaming"), you
	//    can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type
	//    cannot have a fractional DPU allocation.
	MaxCapacity *float64 `json:"maxCapacity,omitempty"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries *int64 `json:"maxRetries,omitempty"`
	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments map[string]*string `json:"nonOverridableArguments,omitempty"`
	// Specifies configuration properties of a job notification.
	NotificationProperty *NotificationProperty `json:"notificationProperty,omitempty"`
	// The number of workers of a defined workerType that are allocated when a job
	// runs.
	//
	// The maximum number of workers you can define are 299 for G.1X, and 149 for
	// G.2X.
	NumberOfWorkers *int64 `json:"numberOfWorkers,omitempty"`
	// The name of the SecurityConfiguration structure to be used with this job.
	SecurityConfiguration *string `json:"securityConfiguration,omitempty"`
	// The tags to use with this job. You may use tags to limit access to the job.
	// For more information about tags in AWS Glue, see AWS Tags in AWS Glue (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html)
	// in the developer guide.
	Tags map[string]*string `json:"tags,omitempty"`
	// The job timeout in minutes. This is the maximum time that a job run can consume
	// resources before it is terminated and enters TIMEOUT status. The default
	// is 2,880 minutes (48 hours).
	Timeout *int64 `json:"timeout,omitempty"`
	// The type of predefined worker that is allocated when a job runs. Accepts
	// a value of Standard, G.1X, or G.2X.
	//
	//    * For the Standard worker type, each worker provides 4 vCPU, 16 GB of
	//    memory and a 50GB disk, and 2 executors per worker.
	//
	//    * For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of
	//    memory, 64 GB disk), and provides 1 executor per worker. We recommend
	//    this worker type for memory-intensive jobs.
	//
	//    * For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of
	//    memory, 128 GB disk), and provides 1 executor per worker. We recommend
	//    this worker type for memory-intensive jobs.
	WorkerType          *string `json:"workerType,omitempty"`
	CustomJobParameters `json:",inline"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       JobParameters `json:"forProvider"`
}

// JobObservation defines the observed state of Job
type JobObservation struct {
	// The unique name that was provided for this job definition.
	Name *string `json:"name,omitempty"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec"`
	Status            JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	JobKind             = "Job"
	JobGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobKind}.String()
	JobKindAPIVersion   = JobKind + "." + GroupVersion.String()
	JobGroupVersionKind = GroupVersion.WithKind(JobKind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
