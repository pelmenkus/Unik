
public class Parser {
    private String[] arr;

    private String slovar="< , template int typename >";
    private int pos, len,in;
    private String[] res;
    public Parser(String s){
        s+=" ";
        System.out.println(s);
        this.arr=s.split(" ");
        this.pos=0;
        this.len=arr.length;
        this.res=new String[len];
        in=0;
    }
    public String parse(){
        //Tm();
        if (!Tm()){
            System.out.println("syntax error at ("+pos+")");
        }
        String s="";
        for(int i=0;i<in;i++){
            s+=res[i]+" ";
        }
        return s;
    }
    private boolean Tm(){
        if (pos<len){
            res[in]="Tm";
            in++;
            if (arr[pos].equals("\n"))
                pos++;
            /*for (int el=0; el<arr.length; el++)
                System.out.println(arr[el]);*/
            //System.out.println(pos);
            if (arr[pos].equals("template")){
                pos++;
                //System.out.println(pos);
                if (arr[pos].equals("<")){
                    pos++;
                    while (pos < len) {
                        if (arr[pos].equals("\n"))
                            pos++;
                        if (arr[pos].equals("template")) {
                            pos++;
                            if (!Tm())
                                return false;
                        }
                            if (arr[pos].equals( "typename")) {
                                pos++;
                                if(!Tpnm())
                                    return false;
                            }
                                if (arr[pos].equals("class")) {
                                    pos++;
                                    if(!Class())
                                        return false;
                                }
                                    if (arr[pos].equals("int")) {
                                        res[in] = "IDENT";
                                        in++;
                                        pos++;
                                        if(!Term())
                                            return false;
                                    }
                                        pos++;
                                    }




                    }
               if (arr[pos-1].equals(">")) {
                   return true;
               }
                else
                    return false;
            }
        }
        return true;
    }

    private boolean Class(){
        //System.out.println(pos);
        res[in]="Class";
        in++;
        if (Term())
            return true;
        else
            return false;
    }
    private boolean Tpnm(){
        //System.out.println(pos);
        res[in]="Typename";
        in++;
        if (Term())
            return true;
        else
            return false;
    }
    private boolean Term(){
        //System.out.println(pos);
        if (slovar.indexOf(arr[pos])==-1) {
            res[in]="E";
            in++;
            pos++;
            return true;
        }
        else
            return false;
    }
}
