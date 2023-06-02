#version 330 core
layout (location = 0) in vec2 aPos;

uniform float position;

void main()
{
    gl_Position = vec4(aPos.x + position, aPos.y, 0.0, 1.0);
}
