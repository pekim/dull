package dull

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var vertexShaderSource = `
	#version 330 core

	layout (location = 0) in vec2 position;
	layout (location = 1) in vec2 texCoords;
	layout (location = 2) in vec4 color;

	out vec2 TexCoords;
	out vec4 Color;

	void main()
	{
		vec2 projection = vec2(1.0, -1.0); // Invert y-axis

		gl_Position = vec4(projection * position, 0.0, 1.0);
		TexCoords = texCoords;
		Color = color;
	}
`

var fragmentShaderSource = `
	#version 330 core

	in vec2 TexCoords;
	in vec4 Color;

	out vec4 color;

	uniform sampler2D textur;

	void main()
	{
		vec4 gamma = vec4(1.0, 1.0, 1.0, 1/1.8);
		vec4 inv_gamma = vec4(1.0, 1.0, 1.0, 1 / gamma.a);

		vec4 colour_linear = pow(Color, inv_gamma);

		vec4 sampled = vec4(1.0, 1.0, 1.0, texture(textur, TexCoords).r);
		vec4 sampled_linear = pow(sampled, inv_gamma);

		vec4 blended = colour_linear * sampled_linear;
		color = pow(blended, vec4(1.0, 1.0, 1.0, gamma));
	}
`

func newProgram() (uint32, error) {
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

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
