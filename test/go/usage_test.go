package test

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	optionsv1 "github.com/omcrgnt/proto/gen/go/options/v1"

	_ "github.com/omcrgnt/proto/gen/go/common/v1"
	_ "github.com/omcrgnt/proto/gen/go/http/v1"
	_ "github.com/omcrgnt/proto/gen/go/logger/v1"
)

func TestAllProtoMessages_UsageOptionMandatory(t *testing.T) {
	// Фильтр: проверяем только сообщения из твоего домена
	const myDomain = "github.com/omcrgnt/proto/gen/go"

	// Проходим по всем зарегистрированным типам сообщений
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		desc := mt.Descriptor()
		fullPath := string(desc.ParentFile().Package())

		// Проверяем только свои пакеты (common.v1, logger.v1 и т.д.)
		// И игнорируем саму опцию и системные типы google/buf
		if strings.HasPrefix(fullPath, "google.") ||
			strings.HasPrefix(fullPath, "buf.") ||
			desc.Name() == "Usage" {
			return true
		}

		t.Run(string(desc.FullName()), func(t *testing.T) {
			fields := desc.Fields()
			for i := 0; i < fields.Len(); i++ {
				field := fields.Get(i)

				// Проверяем наличие опции usage
				opts := field.Options()
				if !proto.HasExtension(opts, optionsv1.E_Usage) {
					t.Errorf("поле '%s' в сообщении '%s' не имеет опции (options.v1.usage)",
						field.Name(), desc.Name())
					continue
				}

				// Проверяем, что значение не пустое
				val := proto.GetExtension(opts, optionsv1.E_Usage).(string)
				if strings.TrimSpace(val) == "" {
					t.Errorf("опция 'usage' для поля '%s' в сообщении '%s' пуста",
						field.Name(), desc.Name())
				}
			}
		})
		return true
	})
}
