package anonymizer

import (
	"testing"

	"github.com/logmanager-oss/logveil/internal/config"
	"github.com/logmanager-oss/logveil/internal/proof"
	"github.com/stretchr/testify/assert"
)

func TestAnonimizer_AnonymizeData(t *testing.T) {
	tests := []struct {
		name               string
		anonymizingDataDir string
		input              map[string]string
		expectedOutput     string
	}{
		{
			name:               "Test AnonymizeData",
			anonymizingDataDir: "../../tests/data/anonymization_data",
			input:              map[string]string{"@timestamp": "2024-06-05T14:59:27.000+00:00", "src_ip": "10.10.10.1", "username": "miloslav.illes", "organization": "Microsoft", "raw": "2024-06-05T14:59:27.000+00:00, 10.10.10.1, miloslav.illes, Microsoft"},
			expectedOutput:     "2024-06-05T14:59:27.000+00:00, 10.20.0.53, ladislav.dosek, Apple",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anonymizer, err := CreateAnonymizer(&config.Config{AnonymizationDataPath: tt.anonymizingDataDir}, &proof.ProofWriter{IsEnabled: false})
			if err != nil {
				t.Fatal(err)
			}
			// Disabling randomization so we know which values to expect
			anonymizer.SetRandFunc(func(int) int { return 1 })
			output := anonymizer.Anonymize(tt.input)

			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
