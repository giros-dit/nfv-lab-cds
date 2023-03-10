package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeFlowExpandedSFlow(t *testing.T) {
	msg := BaseMessage{
		Src:     []byte{},
		Port:    1,
		Payload: getExpandedSFlowDecode(),
	}

	s := &StateSFlow{}

	assert.Nil(t, s.DecodeFlow(msg))
}

func getExpandedSFlowDecode() []byte {
	return []byte{
		0, 0, 0, 5, 0, 0, 0, 1, 1, 2, 3, 4, 0, 0, 0, 0, 5, 167, 139, 219, 5, 118,
		138, 184, 0, 0, 0, 6, 0, 0, 0, 3, 0, 0, 0, 220, 2, 144, 194, 214, 0, 0, 0, 0,
		0, 5, 6, 164, 0, 0, 3, 255, 6, 6, 189, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5,
		6, 164, 0, 0, 0, 0, 0, 5, 6, 171, 0, 0, 0, 2, 0, 0, 3, 233, 0, 0, 0, 6,
		0, 0, 5, 7, 0, 0, 0, 0, 0, 0, 5, 7, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
		0, 144, 0, 0, 0, 1, 0, 0, 5, 234, 0, 0, 0, 4, 0, 0, 0, 128, 8, 6, 168, 250,
		146, 253, 116, 131, 239, 8, 101, 183, 129, 0, 5, 7, 8, 0, 9, 0, 5, 212, 0, 2, 4, 0,
		3, 6, 252, 8, 9, 187, 169, 1, 4, 7, 186, 201, 1, 187, 249, 6, 160, 7, 5, 240, 6, 4,
		4, 0, 0, 6, 0, 123, 119, 210, 0, 0, 165, 105, 7, 171, 145, 234, 102, 0, 252, 187, 162, 227,
		104, 188, 126, 232, 156, 164, 2, 115, 6, 100, 0, 185, 6, 4, 119, 5, 213, 1, 215, 208, 8, 4,
		118, 183, 241, 225, 130, 186, 2, 250, 220, 153, 189, 3, 4, 4, 1, 8, 210, 119, 172, 9, 164, 233,
		1, 8, 171, 226, 196, 195, 3, 152, 9, 5, 6, 181, 4, 7, 0, 0, 0, 3, 0, 0, 0, 220,
		9, 107, 215, 156, 0, 0, 0, 0, 0, 5, 6, 165, 0, 0, 3, 255, 226, 123, 0, 100, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 6, 165, 0, 0, 0, 0, 0, 5, 6, 164, 0, 0, 0, 2,
		0, 0, 3, 233, 0, 0, 0, 6, 0, 0, 3, 184, 0, 0, 0, 0, 0, 0, 3, 184, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0, 144, 0, 0, 0, 1, 0, 0, 5, 190, 0, 0, 0, 4,
		0, 0, 0, 128, 116, 131, 239, 8, 101, 183, 144, 226, 186, 134, 8, 1, 129, 0, 3, 184, 8, 0,
		9, 0, 5, 168, 7, 127, 4, 0, 4, 6, 163, 211, 185, 9, 220, 7, 0, 254, 3, 8, 0, 9,
		130, 136, 179, 1, 2, 2, 7, 5, 250, 4, 128, 6, 0, 1, 7, 1, 0, 0, 1, 1, 8, 0,
		6, 9, 250, 9, 4, 113, 121, 4, 160, 125, 0, 4, 9, 209, 241, 194, 190, 148, 161, 186, 6, 192,
		246, 190, 170, 2, 238, 190, 128, 221, 223, 1, 218, 225, 3, 9, 7, 226, 220, 231, 127, 3, 3, 252,
		7, 9, 161, 247, 218, 8, 8, 174, 133, 4, 213, 245, 149, 218, 5, 4, 200, 128, 139, 5, 0, 115,
		0, 0, 0, 3, 0, 0, 0, 220, 2, 144, 194, 215, 0, 0, 0, 0, 0, 5, 6, 164, 0, 0,
		3, 255, 6, 6, 253, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 6, 164, 0, 0, 0, 0,
		0, 5, 6, 171, 0, 0, 0, 2, 0, 0, 3, 233, 0, 0, 0, 6, 0, 0, 0, 104, 0, 0,
		0, 0, 0, 0, 0, 104, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 144, 0, 0, 0, 1,
		0, 0, 5, 242, 0, 0, 0, 4, 0, 0, 0, 128, 116, 131, 239, 7, 9, 1, 116, 131, 239, 8,
		101, 183, 129, 0, 0, 104, 8, 0, 9, 0, 5, 220, 152, 143, 4, 0, 1, 6, 5, 179, 9, 187,
		191, 101, 190, 2, 144, 182, 0, 0, 130, 4, 252, 4, 160, 192, 138, 8, 219, 124, 128, 6, 0, 235,
		180, 213, 0, 0, 1, 1, 8, 0, 9, 124, 6, 1, 9, 1, 252, 3, 194, 8, 195, 209, 115, 1,
		5, 152, 204, 2, 6, 4, 1, 119, 254, 9, 1, 170, 0, 192, 2, 7, 190, 9, 149, 5, 101, 2,
		128, 122, 0, 190, 1, 109, 188, 175, 4, 8, 152, 1, 142, 108, 2, 100, 2, 124, 125, 195, 5, 8,
		233, 126, 7, 4, 243, 4, 3, 153, 0, 0, 0, 3, 0, 0, 0, 220, 5, 1, 150, 6, 0, 0,
		0, 0, 0, 5, 6, 167, 0, 0, 3, 255, 6, 5, 105, 220, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 5, 6, 167, 0, 0, 0, 0, 0, 5, 6, 164, 0, 0, 0, 2, 0, 0, 3, 233, 0, 0,
		0, 6, 0, 0, 5, 7, 0, 0, 0, 0, 0, 0, 5, 7, 0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 144, 0, 0, 0, 1, 0, 0, 2, 2, 0, 0, 0, 4, 0, 0, 0, 128, 116, 131,
		239, 8, 101, 183, 152, 3, 130, 1, 196, 153, 129, 0, 5, 7, 8, 0, 9, 0, 2, 0, 0, 0,
		4, 0, 126, 7, 119, 188, 185, 9, 221, 8, 2, 116, 144, 0, 9, 139, 3, 112, 2, 0, 8, 124,
		255, 251, 0, 0, 131, 2, 0, 0, 0, 246, 3, 3, 107, 5, 0, 0, 0, 0, 9, 173, 2, 217,
		6, 248, 0, 0, 9, 173, 2, 217, 8, 248, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 6, 9, 153,
		215, 157, 0, 255, 0, 8, 1, 0, 9, 8, 9, 6, 164, 103, 9, 5, 0, 0, 0, 3, 0, 0,
		0, 152, 5, 201, 2, 175, 0, 0, 0, 0, 0, 5, 6, 5, 0, 0, 3, 255, 1, 8, 9, 1,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 6, 5, 0, 0, 0, 0, 0, 5, 6, 164, 0, 0,
		0, 2, 0, 0, 3, 233, 0, 0, 0, 6, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 3,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 6, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0,
		0, 4, 0, 0, 0, 0, 116, 131, 239, 8, 101, 183, 218, 177, 4, 251, 217, 207, 8, 0, 9, 0,
		0, 8, 0, 0, 0, 0, 9, 7, 8, 161, 106, 3, 109, 6, 185, 9, 220, 215, 0, 123, 9, 184,
		0, 8, 116, 122, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 3, 130, 6,
		0, 0, 0, 3, 0, 0, 0, 220, 2, 144, 194, 216, 0, 0, 0, 0, 0, 5, 6, 164, 0, 0,
		3, 255, 6, 7, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 6, 164, 0, 0, 0, 0,
		0, 5, 6, 165, 0, 0, 0, 2, 0, 0, 3, 233, 0, 0, 0, 6, 0, 0, 3, 202, 0, 0,
		0, 0, 0, 0, 3, 202, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 144, 0, 0, 0, 1,
		0, 0, 5, 242, 0, 0, 0, 4, 0, 0, 0, 128, 144, 226, 186, 135, 4, 241, 116, 131, 239, 8,
		101, 183, 129, 0, 3, 202, 8, 0, 9, 0, 5, 220, 147, 0, 4, 0, 7, 6, 225, 131, 1, 159,
		7, 185, 195, 181, 170, 8, 9, 117, 7, 175, 8, 3, 191, 135, 190, 150, 196, 102, 0, 6, 0, 119,
		116, 113, 0, 0, 201, 244, 240, 206, 2, 117, 4, 139, 8, 4, 240, 223, 247, 123, 6, 0, 239, 0,
		9, 116, 152, 153, 191, 0, 124, 2, 7, 8, 3, 178, 166, 150, 3, 218, 163, 175, 121, 8, 4, 210,
		4, 5, 166, 5, 178, 1, 6, 222, 172, 186, 6, 241, 232, 8, 188, 192, 2, 220, 128, 1, 8, 7,
		194, 130, 220, 5, 2, 0, 158, 195, 0, 4, 3, 2, 160, 158, 157, 2, 102, 3, 7, 3, 0, 0,
		1, 3, 3, 4, 1, 1, 4, 2, 187, 255, 188, 3, 4, 138, 9, 180, 104, 233, 212, 239, 123, 237,
		112, 8, 133, 129, 152, 138, 7, 195, 8, 171, 237, 3, 4, 223, 116, 214, 151, 9, 151, 102, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0,
	}
}
