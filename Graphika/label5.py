from OpenGL.GL import *
from OpenGL.GLUT import *
from OpenGL.GLU import *
import random

vertices1 = [(100, 100), (200, 100), (200, 200), (100, 200)]
vertices2 = [(150, 150), (250, 150), (250, 250), (150, 250)]

def reshape(width, height):
    glViewport(0, 0, width, height)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glOrtho(0.0, width, 0.0, height, -1.0, 1.0)
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()

def special_keys(key, x, y):
    global vertices1, vertices2
    step = 10
    if key == GLUT_KEY_LEFT:
        vertices1 = [(x - step, y) for x, y in vertices1]
        vertices2 = [(x - step, y) for x, y in vertices2]
    elif key == GLUT_KEY_RIGHT:
        vertices1 = [(x + step, y) for x, y in vertices1]
        vertices2 = [(x + step, y) for x, y in vertices2]
    elif key == GLUT_KEY_UP:
        vertices1 = [(x, y + step) for x, y in vertices1]
        vertices2 = [(x, y + step) for x, y in vertices2]
    elif key == GLUT_KEY_DOWN:
        vertices1 = [(x, y - step) for x, y in vertices1]
        vertices2 = [(x, y - step) for x, y in vertices2]
    elif key == GLUT_KEY_F1:
        glColor3f(1.0, 0.0, 0.0)
    elif key == GLUT_KEY_F2:
        glColor3f(0.0, 1.0, 0.0)
    elif key == GLUT_KEY_F3:
        glColor3f(0.0, 0.0, 1.0)
    elif key == GLUT_KEY_F4:
        glColor3f(1.0, 1.0, 0.0)
    glutPostRedisplay()

def display():
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glLoadIdentity()

    # Обновляем вершины многоугольников перед объединением
    subject_polygon = vertices1
    clip_polygon = vertices2

    # Отрисовываем первый многоугольник
    glColor3f(0.3, 0.7, 0.8)
    glBegin(GL_POLYGON)
    for vertex in subject_polygon:
        glVertex2f(*vertex)
    glEnd()

    # Отрисовываем второй многоугольник
    glColor3f(0.8, 0.3, 0.7)
    glBegin(GL_POLYGON)
    for vertex in clip_polygon:
        glVertex2f(*vertex)
    glEnd()

    # Объединяем многоугольники и закрашиваем объединенный многоугольник
    combined_polygon = weiler_atherton(subject_polygon, clip_polygon)
    fill_polygon(combined_polygon, (1.0, 1.0, 1.0))

    glutSwapBuffers()

def weiler_atherton(subject_polygon, clip_polygon):
    entry_points = []
    exit_points = []
    intersections = find_intersections(subject_polygon, clip_polygon)
    subject_with_intersections = add_intersections(subject_polygon, intersections)
    clip_with_intersections = add_intersections(clip_polygon, intersections)
    combined_polygon = []
    current_polygon = None
    current_polygon_type = None
    while subject_with_intersections or clip_with_intersections:
        if current_polygon is None:
            if subject_with_intersections:
                current_polygon = subject_with_intersections
                current_polygon_type = "subject"
            elif clip_with_intersections:
                current_polygon = clip_with_intersections
                current_polygon_type = "clip"
        if current_polygon_type == "subject":
            if entry_points:
                if current_polygon:
                    current_point = current_polygon.pop(0)
                    entry_points.append(current_point)
                    if current_point in intersections:
                        current_polygon_type = "clip"
                else:
                    break
            else:
                if current_polygon:
                    current_point = current_polygon.pop(0)
                    entry_points.append(current_point)
                    if current_point in intersections:
                        current_polygon_type = "clip"
        elif current_polygon_type == "clip":
            if exit_points:
                if current_polygon:
                    current_point = current_polygon.pop(0)
                    exit_points.append(current_point)
                    if current_point in intersections:
                        current_polygon_type = "subject"
                else:
                    break
            else:
                if current_polygon:
                    current_point = current_polygon.pop(0)
                    exit_points.append(current_point)
                    if current_point in intersections:
                        current_polygon_type = "subject"
        if entry_points and exit_points and entry_points[0] == exit_points[-1]:
            combined_polygon.extend(merge_polygons(entry_points, exit_points))
            entry_points = []
            exit_points = []
            current_polygon = None
            current_polygon_type = None
    return combined_polygon

def fill_polygon(polygon, color):
    glColor3f(*color)
    glBegin(GL_POLYGON)
    for vertex in polygon:
        glVertex2f(*vertex)
    glEnd()

def main():
    glutInit()
    glutInitDisplayMode(GLUT_RGBA | GLUT_DOUBLE | GLUT_DEPTH)
    glutInitWindowSize(640, 640)
    glutCreateWindow(b"Moving Polygons with PyOpenGL")
    glutDisplayFunc(display)
    glutReshapeFunc(reshape)
    glutSpecialFunc(special_keys)
    glutMainLoop()

if __name__ == "__main__":
    main()
