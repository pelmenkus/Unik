from OpenGL.GL import *
from OpenGL.GLUT import *
from OpenGL.GLU import *

rotation_angle = 0.0
rotation_enabled = True

def square():
    glBegin(GL_QUADS)
    glVertex2f(50, 50)
    glVertex2f(450, 50)
    glVertex2f(450, 450)
    glVertex2f(50, 450)
    glEnd()

def keyPressed(key, x, y):
    global rotation_enabled

    if key == b" ":
        rotation_enabled = not rotation_enabled

def iterate():
    glViewport(0, 0, 500, 500)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glOrtho(0.0, 500, 0.0, 500, 0.0, 1.0)
    glMatrixMode (GL_MODELVIEW)
    glLoadIdentity()

def showScreen():
    global rotation_angle
    global rotation_enabled

    if rotation_enabled:
        rotation_angle += 0.5

    glClearColor(255,255,255, 0)
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glLoadIdentity()
    iterate()
    glColor3f(0, 255, 0)
    glTranslatef(250, 250, 0.0)
    glRotatef(rotation_angle, 0.0, 0.0, 1.0)
    glTranslatef(-250, -250, 0.0)
    square()
    glutSwapBuffers()

glutInit()
glutInitDisplayMode(GLUT_RGBA)
glutInitWindowSize(500, 500)
glutInitWindowPosition(0, 0)
wind = glutCreateWindow(b"Lab1")
glutDisplayFunc(showScreen)
glutIdleFunc(showScreen)
glutKeyboardFunc(keyPressed)
glutMainLoop()