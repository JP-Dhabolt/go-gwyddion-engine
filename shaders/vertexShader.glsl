#version 140
in vec3 position;
in vec4 color;

void main() {
    gl_Position = vec4(position, 1.0);
}
