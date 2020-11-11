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

	// https://chilliant.blogspot.com/2012/08/srgb-approximations-for-hlsl.html
	float to_srgb(float linear) {
		 if (linear <= 0.0031308)
			return linear * 12.92;
		 else
			// Theoretically should be 2.4, but 2.2 results in
			// subjectively nicer text.
			// Reducing any further makes nearby dark colours
			// too similar.
			return 1.055 * pow(linear, 1.0 / 2.2) - 0.055;
	}

	void main()
	{
		vec4 colorLinear = texture(textur, TexCoords);

		// Convert back to sRGB.
		color.r = to_srgb(colorLinear.r);
		color.g = to_srgb(colorLinear.g);
		color.b = to_srgb(colorLinear.b);
		color.a = colorLinear.a;
	}
`
