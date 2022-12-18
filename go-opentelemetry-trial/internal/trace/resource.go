package trace

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

// テレメトリデータがどのサービスからまたはどのサービスインスタンスから来ているかを識別する方法が必要
// テレメトリを生成するエンティティ = Resource
func NewResource(name string, versionKey string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
			semconv.ServiceVersionKey.String(versionKey),
			attribute.String("environment", "demo"),
		),
	)
	return r
}
