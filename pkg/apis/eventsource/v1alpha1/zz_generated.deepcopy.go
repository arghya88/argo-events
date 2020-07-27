// +build !ignore_autogenerated

/*
Copyright 2020 BlackRock, Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"

	common "github.com/argoproj/argo-events/pkg/apis/common"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMQPEventSource) DeepCopyInto(out *AMQPEventSource) {
	*out = *in
	if in.ConnectionBackoff != nil {
		in, out := &in.ConnectionBackoff, &out.ConnectionBackoff
		*out = new(common.Backoff)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMQPEventSource.
func (in *AMQPEventSource) DeepCopy() *AMQPEventSource {
	if in == nil {
		return nil
	}
	out := new(AMQPEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureEventsHubEventSource) DeepCopyInto(out *AzureEventsHubEventSource) {
	*out = *in
	if in.SharedAccessKeyName != nil {
		in, out := &in.SharedAccessKeyName, &out.SharedAccessKeyName
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SharedAccessKey != nil {
		in, out := &in.SharedAccessKey, &out.SharedAccessKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureEventsHubEventSource.
func (in *AzureEventsHubEventSource) DeepCopy() *AzureEventsHubEventSource {
	if in == nil {
		return nil
	}
	out := new(AzureEventsHubEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalendarEventSource) DeepCopyInto(out *CalendarEventSource) {
	*out = *in
	if in.ExclusionDates != nil {
		in, out := &in.ExclusionDates, &out.ExclusionDates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserPayload != nil {
		in, out := &in.UserPayload, &out.UserPayload
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalendarEventSource.
func (in *CalendarEventSource) DeepCopy() *CalendarEventSource {
	if in == nil {
		return nil
	}
	out := new(CalendarEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmitterEventSource) DeepCopyInto(out *EmitterEventSource) {
	*out = *in
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ConnectionBackoff != nil {
		in, out := &in.ConnectionBackoff, &out.ConnectionBackoff
		*out = new(common.Backoff)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmitterEventSource.
func (in *EmitterEventSource) DeepCopy() *EmitterEventSource {
	if in == nil {
		return nil
	}
	out := new(EmitterEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSource) DeepCopyInto(out *EventSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSource.
func (in *EventSource) DeepCopy() *EventSource {
	if in == nil {
		return nil
	}
	out := new(EventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceList) DeepCopyInto(out *EventSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceList.
func (in *EventSourceList) DeepCopy() *EventSourceList {
	if in == nil {
		return nil
	}
	out := new(EventSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceSpec) DeepCopyInto(out *EventSourceSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(Service)
		(*in).DeepCopyInto(*out)
	}
	if in.Replica != nil {
		in, out := &in.Replica, &out.Replica
		*out = new(int32)
		**out = **in
	}
	if in.Minio != nil {
		in, out := &in.Minio, &out.Minio
		*out = make(map[string]common.S3Artifact, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Calendar != nil {
		in, out := &in.Calendar, &out.Calendar
		*out = make(map[string]CalendarEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = make(map[string]FileEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make(map[string]ResourceEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = make(map[string]WebhookContext, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AMQP != nil {
		in, out := &in.AMQP, &out.AMQP
		*out = make(map[string]AMQPEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = make(map[string]KafkaEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.MQTT != nil {
		in, out := &in.MQTT, &out.MQTT
		*out = make(map[string]MQTTEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		*out = make(map[string]NATSEventsSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.SNS != nil {
		in, out := &in.SNS, &out.SNS
		*out = make(map[string]SNSEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.SQS != nil {
		in, out := &in.SQS, &out.SQS
		*out = make(map[string]SQSEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.PubSub != nil {
		in, out := &in.PubSub, &out.PubSub
		*out = make(map[string]PubSubEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = make(map[string]GithubEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = make(map[string]GitlabEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.HDFS != nil {
		in, out := &in.HDFS, &out.HDFS
		*out = make(map[string]HDFSEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = make(map[string]SlackEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.StorageGrid != nil {
		in, out := &in.StorageGrid, &out.StorageGrid
		*out = make(map[string]StorageGridEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.AzureEventsHub != nil {
		in, out := &in.AzureEventsHub, &out.AzureEventsHub
		*out = make(map[string]AzureEventsHubEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Stripe != nil {
		in, out := &in.Stripe, &out.Stripe
		*out = make(map[string]StripeEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Emitter != nil {
		in, out := &in.Emitter, &out.Emitter
		*out = make(map[string]EmitterEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = make(map[string]RedisEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.NSQ != nil {
		in, out := &in.NSQ, &out.NSQ
		*out = make(map[string]NSQEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Pulsar != nil {
		in, out := &in.Pulsar, &out.Pulsar
		*out = make(map[string]PulsarEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Generic != nil {
		in, out := &in.Generic, &out.Generic
		*out = make(map[string]GenericEventSource, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceSpec.
func (in *EventSourceSpec) DeepCopy() *EventSourceSpec {
	if in == nil {
		return nil
	}
	out := new(EventSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceStatus) DeepCopyInto(out *EventSourceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceStatus.
func (in *EventSourceStatus) DeepCopy() *EventSourceStatus {
	if in == nil {
		return nil
	}
	out := new(EventSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileEventSource) DeepCopyInto(out *FileEventSource) {
	*out = *in
	out.WatchPathConfig = in.WatchPathConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileEventSource.
func (in *FileEventSource) DeepCopy() *FileEventSource {
	if in == nil {
		return nil
	}
	out := new(FileEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericEventSource) DeepCopyInto(out *GenericEventSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericEventSource.
func (in *GenericEventSource) DeepCopy() *GenericEventSource {
	if in == nil {
		return nil
	}
	out := new(GenericEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubEventSource) DeepCopyInto(out *GithubEventSource) {
	*out = *in
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookContext)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIToken != nil {
		in, out := &in.APIToken, &out.APIToken
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.WebhookSecret != nil {
		in, out := &in.WebhookSecret, &out.WebhookSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubEventSource.
func (in *GithubEventSource) DeepCopy() *GithubEventSource {
	if in == nil {
		return nil
	}
	out := new(GithubEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabEventSource) DeepCopyInto(out *GitlabEventSource) {
	*out = *in
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookContext)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessToken != nil {
		in, out := &in.AccessToken, &out.AccessToken
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabEventSource.
func (in *GitlabEventSource) DeepCopy() *GitlabEventSource {
	if in == nil {
		return nil
	}
	out := new(GitlabEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HDFSEventSource) DeepCopyInto(out *HDFSEventSource) {
	*out = *in
	out.WatchPathConfig = in.WatchPathConfig
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KrbCCacheSecret != nil {
		in, out := &in.KrbCCacheSecret, &out.KrbCCacheSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.KrbKeytabSecret != nil {
		in, out := &in.KrbKeytabSecret, &out.KrbKeytabSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.KrbConfigConfigMap != nil {
		in, out := &in.KrbConfigConfigMap, &out.KrbConfigConfigMap
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HDFSEventSource.
func (in *HDFSEventSource) DeepCopy() *HDFSEventSource {
	if in == nil {
		return nil
	}
	out := new(HDFSEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaEventSource) DeepCopyInto(out *KafkaEventSource) {
	*out = *in
	if in.ConnectionBackoff != nil {
		in, out := &in.ConnectionBackoff, &out.ConnectionBackoff
		*out = new(common.Backoff)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaEventSource.
func (in *KafkaEventSource) DeepCopy() *KafkaEventSource {
	if in == nil {
		return nil
	}
	out := new(KafkaEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTTEventSource) DeepCopyInto(out *MQTTEventSource) {
	*out = *in
	if in.ConnectionBackoff != nil {
		in, out := &in.ConnectionBackoff, &out.ConnectionBackoff
		*out = new(common.Backoff)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTTEventSource.
func (in *MQTTEventSource) DeepCopy() *MQTTEventSource {
	if in == nil {
		return nil
	}
	out := new(MQTTEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSEventsSource) DeepCopyInto(out *NATSEventsSource) {
	*out = *in
	if in.ConnectionBackoff != nil {
		in, out := &in.ConnectionBackoff, &out.ConnectionBackoff
		*out = new(common.Backoff)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSEventsSource.
func (in *NATSEventsSource) DeepCopy() *NATSEventsSource {
	if in == nil {
		return nil
	}
	out := new(NATSEventsSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSQEventSource) DeepCopyInto(out *NSQEventSource) {
	*out = *in
	if in.ConnectionBackoff != nil {
		in, out := &in.ConnectionBackoff, &out.ConnectionBackoff
		*out = new(common.Backoff)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSQEventSource.
func (in *NSQEventSource) DeepCopy() *NSQEventSource {
	if in == nil {
		return nil
	}
	out := new(NSQEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubEventSource) DeepCopyInto(out *PubSubEventSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubEventSource.
func (in *PubSubEventSource) DeepCopy() *PubSubEventSource {
	if in == nil {
		return nil
	}
	out := new(PubSubEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarEventSource) DeepCopyInto(out *PulsarEventSource) {
	*out = *in
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	if in.ConnectionBackoff != nil {
		in, out := &in.ConnectionBackoff, &out.ConnectionBackoff
		*out = new(common.Backoff)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarEventSource.
func (in *PulsarEventSource) DeepCopy() *PulsarEventSource {
	if in == nil {
		return nil
	}
	out := new(PulsarEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEventSource) DeepCopyInto(out *RedisEventSource) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEventSource.
func (in *RedisEventSource) DeepCopy() *RedisEventSource {
	if in == nil {
		return nil
	}
	out := new(RedisEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEventSource) DeepCopyInto(out *ResourceEventSource) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(ResourceFilter)
		(*in).DeepCopyInto(*out)
	}
	out.GroupVersionResource = in.GroupVersionResource
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]ResourceEventType, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEventSource.
func (in *ResourceEventSource) DeepCopy() *ResourceEventSource {
	if in == nil {
		return nil
	}
	out := new(ResourceEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFilter) DeepCopyInto(out *ResourceFilter) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Selector, len(*in))
		copy(*out, *in)
	}
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]Selector, len(*in))
		copy(*out, *in)
	}
	in.CreatedBy.DeepCopyInto(&out.CreatedBy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFilter.
func (in *ResourceFilter) DeepCopy() *ResourceFilter {
	if in == nil {
		return nil
	}
	out := new(ResourceFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSEventSource) DeepCopyInto(out *SNSEventSource) {
	*out = *in
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookContext)
		**out = **in
	}
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSEventSource.
func (in *SNSEventSource) DeepCopy() *SNSEventSource {
	if in == nil {
		return nil
	}
	out := new(SNSEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQSEventSource) DeepCopyInto(out *SQSEventSource) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQSEventSource.
func (in *SQSEventSource) DeepCopy() *SQSEventSource {
	if in == nil {
		return nil
	}
	out := new(SQSEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackEventSource) DeepCopyInto(out *SlackEventSource) {
	*out = *in
	if in.SigningSecret != nil {
		in, out := &in.SigningSecret, &out.SigningSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookContext)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackEventSource.
func (in *SlackEventSource) DeepCopy() *SlackEventSource {
	if in == nil {
		return nil
	}
	out := new(SlackEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageGridEventSource) DeepCopyInto(out *StorageGridEventSource) {
	*out = *in
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookContext)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(StorageGridFilter)
		**out = **in
	}
	if in.AuthToken != nil {
		in, out := &in.AuthToken, &out.AuthToken
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageGridEventSource.
func (in *StorageGridEventSource) DeepCopy() *StorageGridEventSource {
	if in == nil {
		return nil
	}
	out := new(StorageGridEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageGridFilter) DeepCopyInto(out *StorageGridFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageGridFilter.
func (in *StorageGridFilter) DeepCopy() *StorageGridFilter {
	if in == nil {
		return nil
	}
	out := new(StorageGridFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StripeEventSource) DeepCopyInto(out *StripeEventSource) {
	*out = *in
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookContext)
		**out = **in
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.EventFilter != nil {
		in, out := &in.EventFilter, &out.EventFilter
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StripeEventSource.
func (in *StripeEventSource) DeepCopy() *StripeEventSource {
	if in == nil {
		return nil
	}
	out := new(StripeEventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(v1.Container)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatchPathConfig) DeepCopyInto(out *WatchPathConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatchPathConfig.
func (in *WatchPathConfig) DeepCopy() *WatchPathConfig {
	if in == nil {
		return nil
	}
	out := new(WatchPathConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookContext) DeepCopyInto(out *WebhookContext) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookContext.
func (in *WebhookContext) DeepCopy() *WebhookContext {
	if in == nil {
		return nil
	}
	out := new(WebhookContext)
	in.DeepCopyInto(out)
	return out
}
