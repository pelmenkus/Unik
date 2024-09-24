#!/usr/bin/env python3

import glfw
from OpenGL.GL import *
from math import cos, sin
import numpy as np

a = 0.0
b = 0.0
size = 0.7
fill = True
def main():
    if not glfw.init():
        return
    window = glfw.create_window(700, 700, "CUBE", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)

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

    # move
    glMultMatrixf([1, 0, 0, 0,
                   0, 1, 0, 0,
                   0, 0, 1, 0,
                   0.75, 0.75, 0, 1])


    def projection():
        a_rad = np.radians(a)
        b_rad = np.radians(b)

        rotate_y = np.array([
            [cos(a_rad), 0, sin(a_rad), 0],
            [0, 1, 0, 0],
            [-sin(a_rad), 0, cos(a_rad), 0],
            [0, 0, 0, 1]
        ])

        rotate_x = np.array([
            [1, 0, 0, 0],
            [0, cos(b_rad), sin(b_rad), 0],
            [0, sin(b_rad), cos(b_rad), 0],
            [0, 0, 0, 1]
        ])

        tree_points_projection = np.array([
            [0.87, 0, 0.5, 0],
            [-0.09, 0.98, 0.15, 0],
            [0.98, 0.35, -1.7, 1],
            [0.49, 0.17, -0.85, 2]
        ])

        glMultMatrixf(rotate_x)
        glMultMatrixf(rotate_y)
        glMultMatrixf(tree_points_projection)

    projection()

    def cube(x):
        x = x/2
        glBegin(GL_QUADS)
        glColor3f(1.0, 0.0, 1.0);
        glVertex3f(-x, -x, -x)
        glVertex3f(-x,  x, -x)
        glVertex3f(-x,  x,  x)
        glVertex3f(-x, -x,  x)
        glColor3f(1.0, 0.5, 0.5);
        glVertex3f( x, -x, -x)
        glVertex3f( x, -x,  x)
        glVertex3f( x,  x,  x)
        glVertex3f( x,  x, -x)
        glColor3f(0.5, 1.0, 0.0);
        glVertex3f(-x, -x, -x)
        glVertex3f(-x, -x,  x)
        glVertex3f( x, -x,  x)
        glVertex3f( x, -x, -x)
        glColor3f(1.0, 1.0, 0.0);
        glVertex3f(-x, x, -x)
        glVertex3f(-x, x,  x)
        glVertex3f( x, x,  x)
        glVertex3f( x, x, -x)
        glColor3f(0.0, 1.0, 1.0);
        glVertex3f(-x, -x, -x)
        glVertex3f( x, -x, -x)
        glVertex3f( x,  x, -x)
        glVertex3f(-x,  x, -x)
        glColor3f(1.0, 0.5, 0.0);
        glVertex3f(-x, -x,  x)
        glVertex3f( x, -x,  x)
        glVertex3f( x,  x,  x)
        glVertex3f(-x,  x,  x)
        glEnd()

    cube(0.1)

    glLoadIdentity()

    x = 0.7
    y = 0.7
    projection()

    cube(size)

    glfw.swap_buffers(window)
    glfw.poll_events()

def key_callback(window, key, scancode, action, mods):
    global a
    global b
    if action == glfw.PRESS or action == glfw.REPEAT:
        if key == glfw.KEY_RIGHT:
            b += 0.5
        elif key == glfw.KEY_LEFT:
            b -= 0.5
        elif key == glfw.KEY_UP:
            a += 0.5
        elif key == glfw.KEY_DOWN:
            a -= 0.5
        elif key == glfw.KEY_F:
            global fill
            fill = not fill
            if fill:
                glPolygonMode(GL_FRONT_AND_BACK, GL_FILL);
            else:
                glPolygonMode(GL_FRONT_AND_BACK, GL_LINE);



if __name__ == "__main__":
    main()
