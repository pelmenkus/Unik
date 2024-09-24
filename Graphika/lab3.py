import random
from math import cos as cos
from math import sin as sin

import glfw
import numpy as np
from OpenGL.GL import *

alpha = 0.0
beta = 0.0
size = 0.7
fill = True


def main():
    if not glfw.init():
        return
    window = glfw.create_window(640, 640, "LAB 3", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    glfw.set_scroll_callback(window, scroll_callback)

    glEnable(GL_DEPTH_TEST)
    glDepthFunc(GL_LESS)
    while not glfw.window_should_close(window):
        display(window)
    glfw.destroy_window(window)
    glfw.terminate()


def display(window):
    glLoadIdentity()
    glClear(GL_COLOR_BUFFER_BIT)
    glClear(GL_DEPTH_BUFFER_BIT)
    glMatrixMode(GL_PROJECTION)
    global alpha
    global beta

    glMultMatrixf([1, 0, 0, 0,
                   0, 1, 0, 0,
                   0, 0, 1, 0,
                   0.75, 0.75, 0, 1])

    theta = 0.7
    phi = 0.8

    glMultMatrixf([
        cos(phi), 0, sin(phi), 0,
        sin(phi) * cos(theta), cos(theta), -cos(phi) * sin(theta), 0,
        sin(phi) * sin(theta), -sin(theta), cos(phi) * cos(theta), 0,
        0, 0, 0, 1,
    ])

    def projection():
        alpha_rad = np.radians(alpha)
        beta_rad = np.radians(beta)

        rotate_y = np.array([
            [cos(alpha_rad), 0, sin(alpha_rad), 0],
            [0, 1, 0, 0],
            [-sin(alpha_rad), 0, cos(alpha_rad), 0],
            [0, 0, 0, 1]
        ])

        rotate_x = np.array([
            [1, 0, 0, 0],
            [0, cos(beta_rad), -sin(beta_rad), 0],
            [0, sin(beta_rad), cos(beta_rad), 0],
            [0, 0, 0, 1]
        ])


        glMultMatrixf(rotate_y)
        glMultMatrixf(rotate_x)

    def cylinder(radius, height, slices=30):
        angle = (2 * np.pi) / slices
        glBegin(GL_QUAD_STRIP)
        for i in range(slices + 1):
            glColor3f(random.randint(0,10)/10, random.randint(0,10)/10, random.randint(0,10)/10)
            x = radius * np.cos(i * angle)
            y = radius * np.sin(i * angle)
            glVertex3f(x, y, 0)
            glVertex3f(x, y, height)
        glEnd()

        glColor3f(0.3, 0.7, 0.8)
        glBegin(GL_POLYGON)
        for i in range(slices):
            x = radius * np.cos(i * angle)
            y = radius * np.sin(i * angle)
            glVertex3f(x, y, 0)
        glEnd()

        glBegin(GL_POLYGON)
        for i in range(slices):
            x = radius * np.cos(i * angle)
            y = radius * np.sin(i * angle)
            glVertex3f(x, y, height)
        glEnd()

    cylinder(0.1, 0.2)

    glLoadIdentity()

    x = 0.7
    y = 0.7
    fz = np.sqrt(x * x + y * y)
    projection()

    cylinder(size * 0.5, size)

    glfw.swap_buffers(window)
    glfw.poll_events()


def key_callback(window, key, scancode, action, mods):
    global alpha
    global beta
    if action == glfw.PRESS or action == glfw.REPEAT:
        if key == glfw.KEY_RIGHT:
            alpha += 1
        elif key == glfw.KEY_LEFT:
            alpha -= 1
        elif key == glfw.KEY_UP:
            beta -= 1
        elif key == glfw.KEY_DOWN:
            beta += 1
        elif key == glfw.KEY_F:
            global fill
            fill = not fill
            if fill:
                glPolygonMode(GL_FRONT_AND_BACK, GL_FILL)
            else:
                glPolygonMode(GL_FRONT_AND_BACK, GL_LINE)


def scroll_callback(window, xoffset, yoffset):
    global size

    if xoffset > 0:
        size -= yoffset / 10
    else:
        size += yoffset / 10


if __name__ == "__main__":
    main()