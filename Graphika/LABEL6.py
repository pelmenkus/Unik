from OpenGL.GL import *
from OpenGL.GLUT import *
from OpenGL.GLU import *
import glfw
from math import *
from PIL import Image

up, down, left, right, show_faces, size, density, texture_enabled, texture_id, texture_filename = False, False, False, False, True, 0.8, 20, True, None, 'ryan.bmp'
position, velocity, prev_velocity = [0.0, 0.0, 0.0], [0.005, 0.008, 0.003], [0.0, 0.0, 0.0]
axes = [
    [1.0, 0.0, 0.0],
    [0.0, -1.0, 0.0],
    [0.0, 0.0, 1.0],
]


def generate_sphere():
    def generate_circle(axis1):
        def get_coorinates(axis1, axis2, axis3):
            l = []
            for i in range(3):
                a = 0
                a += axis1 * axes[0][i]
                a += axis2 * axes[1][i]
                a += axis3 * axes[2][i]
                a *= size
                a += position[i]
                l.append(a)
            return l

        l = []
        for i in range(1, density + 1):
            angle = radians(i * 360 / density)
            radius = sqrt(1 - (axis1 ** 2))
            l.append(get_coorinates(axis1, radius * sin(angle), radius * cos(angle)))
        return l

    sphere = []
    delta = 2.0 / density
    for i in range(density + 1):
        axis1 = 1.0 - delta * i
        sphere.append(generate_circle(axis1))
    return sphere


def init():
    glEnable(GL_DEPTH_TEST)
    glEnable(GL_LIGHTING)
    glEnable(GL_LIGHT0)
    glEnable(GL_COLOR_MATERIAL)
    glLightfv(GL_LIGHT0, GL_POSITION, [0.5, 0.0, 0.0, 0.0])
    glLightfv(GL_LIGHT0, GL_DIFFUSE, [1.0, 1.0, 1.0, 1.0])
    glLightfv(GL_LIGHT0, GL_SPECULAR, [1.0, 1.0, 1.0, 1.0])


def load_texture(filename):
    image = Image.open(filename).convert('RGB')
    image_data = image.tobytes("raw", "RGBX", 0, -1)

    texture = glGenTextures(1)
    glBindTexture(GL_TEXTURE_2D, texture)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR)
    glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, image.width, image.height, 0, GL_RGBA, GL_UNSIGNED_BYTE, image_data)

    return texture


def main():
    global texture_id, texture_filename
    if not glfw.init():
        return

    window = glfw.create_window(800, 800, "Opengl GLFW Window", None, None)

    if not window:
        glfw.terminate()
        return

    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    glfw.set_scroll_callback(window, scroll_callback)
    init()

    texture_id = load_texture(texture_filename)

    while not glfw.window_should_close(window):
        update()
        glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
        glLoadIdentity()
        glfw.poll_events()
        display()
        rotate()
        glfw.swap_buffers(window)

    glfw.destroy_window(window)
    glfw.terminate()


def display():
    sphere = generate_sphere()
    if texture_enabled:
        glEnable(GL_TEXTURE_2D)
        glBindTexture(GL_TEXTURE_2D, texture_id)
        for i in range(density):
            for j in range(density):
                glBegin(GL_POLYGON)
                glTexCoord2f(i / density, j / density)
                glNormal3fv(sphere[i][j])
                glVertex3f(*sphere[i][j])
                glTexCoord2f((i + 1) / density, j / density)
                glNormal3fv(sphere[i + 1][j])
                glVertex3f(*sphere[i + 1][j])
                glTexCoord2f((i + 1) / density, (j - 1) / density)
                glNormal3fv(sphere[i + 1][j - 1])
                glVertex3f(*sphere[i + 1][j - 1])
                glTexCoord2f(i / density, (j - 1) / density)
                glNormal3fv(sphere[i][j - 1])
                glVertex3f(*sphere[i][j - 1])
                glEnd()
        glDisable(GL_TEXTURE_2D)
    elif show_faces:
        for i in range(density):
            for j in range(density):
                glBegin(GL_POLYGON)
                glColor3f(i / density, j / density, 1.0)
                glNormal3fv(sphere[i][j])
                glVertex3f(*sphere[i][j])
                glVertex3f(*sphere[i + 1][j])
                glVertex3f(*sphere[i + 1][j - 1])
                glVertex3f(*sphere[i][j - 1])
                glEnd()
    else:
        for i in range(density):
            for j in range(density):
                glBegin(GL_LINES)
                glColor3f(1.0, 1.0, 1.0)
                glVertex3f(*sphere[i][j])
                glVertex3f(*sphere[i][j - 1])
                glEnd()
                glBegin(GL_LINES)
                glColor3f(1.0, 1.0, 1.0)
                glVertex3f(*sphere[i][j])
                glVertex3f(*sphere[i + 1][j])
                glEnd()


def key_callback(window, key, scancode, action, mods):
    global up, down, left, right, cube, show_faces, density, texture_enabled, velocity, prev_velocity, position
    if action == glfw.PRESS:
        if key == glfw.KEY_W:
            up = True
        if key == glfw.KEY_A:
            left = True
        if key == glfw.KEY_S:
            down = True
        if key == glfw.KEY_D:
            right = True
        if key == glfw.KEY_SPACE:
            show_faces = not show_faces
            texture_enabled = False
        if key == glfw.KEY_UP:
            density += 1
        if key == glfw.KEY_DOWN and density > 5:
            density -= 1
        if key == glfw.KEY_ENTER:
            texture_enabled = not texture_enabled
        if key == glfw.KEY_BACKSPACE:
            velocity, prev_velocity = prev_velocity, velocity
        if key == glfw.KEY_RIGHT_SHIFT:
            position = [0.0, 0.0, 0.0]
            velocity = [0.0, 0.0, 0.0]
            prev_velocity = [0.005, 0.008, 0.003]

    if action == glfw.RELEASE:
        if key == glfw.KEY_W:
            up = False
        if key == glfw.KEY_A:
            left = False
        if key == glfw.KEY_S:
            down = False
        if key == glfw.KEY_D:
            right = False


def scroll_callback(window, xoffset, yoffset):
    global size, velocity
    tvelocity = velocity
    velocity = [0.0, 0.0, 0.0]
    if size > 1.0:
        size = 1.0
    elif size < 0.1:
        size = 0.1
    elif size >= 0.1 and yoffset > 0:
        size -= yoffset / 100
    elif size <= 1.0 and yoffset < 0:
        size -= yoffset / 100
    velocity = tvelocity


def rotate():
    global up, down, left, right, axes
    for i in range(3):
        r_sphere = 1
        if up:
            angle = degrees(atan2(axes[i][2], axes[i][1]) + 2 * pi)
            angle -= 1
            angle = radians(angle)
            r = sqrt(r_sphere - axes[i][0] ** 2)
            axes[i][1] = r * cos(angle)
            axes[i][2] = r * sin(angle)
        elif down:
            angle = degrees(atan2(axes[i][2], axes[i][1]) + 2 * pi)
            angle += 1
            angle = radians(angle)
            r = sqrt(r_sphere - axes[i][0] ** 2)
            axes[i][1] = r * cos(angle)
            axes[i][2] = r * sin(angle)
        if left:
            angle = degrees(atan2(axes[i][2], axes[i][0]) + 2 * pi)
            angle += 1
            angle = radians(angle)
            r = sqrt(r_sphere - axes[i][1] ** 2)
            axes[i][0] = r * cos(angle)
            axes[i][2] = r * sin(angle)
        elif right:
            angle = degrees(atan2(axes[i][2], axes[i][0]) + 2 * pi)
            angle -= 1
            angle = radians(angle)
            r = sqrt(r_sphere - axes[i][1] ** 2)
            axes[i][0] = r * cos(angle)
            axes[i][2] = r * sin(angle)


def update():
    global position, velocity, axes
    for i in range(3):
        position[i] += velocity[i]
    for i in range(3):
        if abs(position[i]) + size >= 1.0:
            velocity[i] *= -1


if __name__ == "__main__":
    main()
