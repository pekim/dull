package gl

var gammaVertexShaderSource = `
	#version 330 core

	layout (location = 0) in vec2 position;
	layout (location = 1) in vec2 texCoords;

	out vec2 TexCoords;

	void main()
	{
		gl_Position = vec4(position.x, position.y, 0.0, 1.0);
		TexCoords = texCoords;
	}
`

var gammaFragmentShaderSource = `
	#version 330 core

	in vec2 TexCoords;

	out vec4 color;

	uniform sampler2D textur;
	uniform float gamma;

	void main()
	{
		color = texture(textur, TexCoords);

		// Apply gamma correction.
		color = pow(color, vec4(vec3(1 / gamma), 1.0));
	}
`
