package gl

var renderVertexShaderSource = `
	#version 330 core

	layout (location = 0) in vec2 position;
	layout (location = 1) in vec2 texCoords;
	layout (location = 2) in vec4 color;

	out vec2 TexCoords;
	out vec4 Color;

	// https://chilliant.blogspot.com/2012/08/srgb-approximations-for-hlsl.html
	float to_linear(float srgb) {
		if (srgb <= 0.04045)
			return srgb / 12.92;
		else
			return pow((srgb + 0.055) / 1.055, 2.4);
	}

	void main()
	{
		vec2 projection = vec2(1.0, -1.0); // Invert y-axis

		gl_Position = vec4(projection * position, 0.0, 1.0);
		TexCoords = texCoords;

		// Convert from sRGB to linear.
		Color.r = to_linear(color.r);
		Color.g = to_linear(color.g);
		Color.b = to_linear(color.b);
		Color.a = color.a;
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
