import javax.swing.*;
import javax.swing.event.ChangeEvent;
import javax.swing.event.ChangeListener;


public class kentavr {
    private JPanel counters;
    private JSpinner nCnt;
    private JLabel nogi1;
    private JSpinner rCnt;
    private JLabel ruki;
    private CanvasPanel Risunok;

    public kentavr(){
        rCnt.setValue(2);
        nCnt.setValue(4);
        nCnt.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int kolvo = (int) nCnt.getValue();
                Risunok.setNogi(kolvo);
            }
        });
        rCnt.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int kolvo = (int) rCnt.getValue();
                Risunok.setRuki(kolvo);
            }
        });
    }
    public static void main(String[] args) {
        JFrame frame = new JFrame("kentavr");
        frame.setContentPane(new kentavr().counters);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);
    }

    private void createUIComponents() {
        // TODO: place custom component creation code here
        Risunok = new CanvasPanel();
    }
}
