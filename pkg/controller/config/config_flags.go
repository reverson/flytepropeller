// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "kube-config"), *new(string), "Path to kubernetes client config file.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "master"), *new(string), "")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "workers"), 2, "Number of threads to process workflows")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "workflow-reeval-duration"), "30s", "Frequency of re-evaluating workflows")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "downstream-eval-duration"), "60s", "Frequency of re-evaluating downstream tasks")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "limit-namespace"), "all", "Namespaces to watch for this propeller")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "prof-port"), "10254", "Profiler port")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metadata-prefix"), *new(string), "MetadataPrefix should be used if all the metadata for Flyte executions should be stored under a specific prefix in CloudStorage. If not specified,  the data will be stored in the base container directly.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.type"), "simple", "Type of composite queue to use for the WorkQueue")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.queue.type"), "default", "Type of RateLimiter to use for the WorkQueue")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.queue.base-delay"), "10s", "base backoff delay for failure")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.queue.max-delay"), "10s", "Max backoff delay for failure")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "queue.queue.rate"), int64(10), "Bucket Refill rate per second")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "queue.queue.capacity"), 100, "Bucket capacity as number of items")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.type"), "default", "Type of RateLimiter to use for the WorkQueue")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.base-delay"), "10s", "base backoff delay for failure")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.max-delay"), "10s", "Max backoff delay for failure")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.rate"), int64(10), "Bucket Refill rate per second")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.capacity"), 100, "Bucket capacity as number of items")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.batching-interval"), "1s", "Duration for which downstream updates are buffered")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "queue.batch-size"), -1, "Number of downstream triggered top-level objects to re-enqueue every duration. -1 indicates all available.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metrics-prefix"), "flyte:", "An optional prefix for all published metrics.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "enable-admin-launcher"), false, " Enable remote Workflow launcher to Admin")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "max-workflow-retries"), 50, "Maximum number of retries per workflow")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "max-ttl-hours"), 23, "Maximum number of hours a completed workflow should be retained. Number between 1-23 hours")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "gc-interval"), "30m", "Run periodic GC every 30 minutes")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "leader-election.enabled"), *new(bool), "Enables/Disables leader election.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.lock-config-map.Namespace"), *new(string), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.lock-config-map.Name"), *new(string), "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.lease-duration"), "15s", "Duration that non-leader candidates will wait to force acquire leadership. This is measured against time of last observed ack.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.renew-deadline"), "10s", "Duration that the acting master will retry refreshing leadership before giving up.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.retry-period"), "2s", "Duration the LeaderElector clients should wait between tries of actions.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "publish-k8s-events"), *new(bool), "Enable events publishing to K8s events API.")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "max-output-size-bytes"), *new(int64), "Maximum size of outputs per task")
	return cmdFlags
}
