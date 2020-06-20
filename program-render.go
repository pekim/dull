package dull

var renderVertexShaderSource = `
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

var renderFragmentShaderSource = `
	#version 330 core

	in vec2 TexCoords;
	in vec4 Color;

	out vec4 color;

	uniform sampler2D textur;

	void main()
	{
		vec4 sampled = vec4(1.0, 1.0, 1.0, texture(textur, TexCoords).r);
		color = Color * sampled;
	}
`
