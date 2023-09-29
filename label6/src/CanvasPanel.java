import javax.swing.*;
import java.awt.*;

public class CanvasPanel extends JPanel {
    private int radius = 100;
    private int nogi = 4;
    private int ruki = 2;
    public void setRadius(int r) {
        radius = r;
        repaint();
    }
    public void setNogi(int r) {
        nogi=r;
        repaint();
    }
    public void setRuki(int r) {
        ruki=r;
        repaint();
    }
    protected void paintComponent(Graphics g) {
        super.paintComponent(g);
        g.setColor(Color.MAGENTA);
        int x=10, y=10;
       // g.drawOval();
        g.fillRect(x+radius/2,y+radius,radius/2,radius*2);
        g.fillRect(x+radius/2,y+2*radius,radius*2,radius);
        g.fillOval(x+20,y+10,radius,radius);
        int razm=radius*2/nogi;
        int space=razm/5;
        int sizex=x+radius/2;
        int sizey=y+2*radius+radius;
        for (int i=0; i<nogi; i++){
            g.setColor(Color.MAGENTA);
            g.drawRect(sizex,sizey,razm-space,radius);
            g.setColor(Color.BLACK);
            g.fillRect(sizex,sizey+radius,razm-space,5);
            sizex=sizex+razm;
        }
        sizex=x;
        sizey=y+2*radius;
        razm=radius/ruki;
        space=razm/5;
        for (int i=0; i<ruki; i++){
            g.setColor(Color.MAGENTA);
            g.drawRect(sizex,sizey,radius/2,razm-space);
            g.setColor(Color.BLACK);
            g.fillRect(sizex-5,sizey,5,razm-space);
            sizey=sizey+razm;
        }
        int[] xPoints={x+radius/2+radius*2-20,x+radius/2+radius*2-10,x+radius/2+radius*2};
        int[] yPoints={y+2*radius,y+2*radius-35,y+2*radius};
        g.fillPolygon(xPoints,yPoints,3);
        g.fillOval(x+radius/3,y+radius/3,radius/10,radius/10);
        g.fillOval(x+radius,y+radius/3,radius/10,radius/10);
        g.fillOval(x+radius/2+5,y+radius/2+10,radius/5,radius/10);
    }
}