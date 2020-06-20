package gl

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func newProgram(vertexShaderSource string, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	programId := gl.CreateProgram()

	gl.AttachShader(programId, vertexShader)
	gl.AttachShader(programId, fragmentShader)
	gl.LinkProgram(programId)

	var status int32
	gl.GetProgramiv(programId, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(programId, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(programId, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return programId, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	csources, free := gl.Strs(source + "\x00")
	defer free()

	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, csources, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile shader\n%v\n%v", source, log)
	}

	return shader, nil
}
