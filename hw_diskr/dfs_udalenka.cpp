#include<bits/stdc++.h>
using namespace std;

void dfs(int& tmp, int& razm, vector<int>& dist, vector<bool>& flag, vector <vector<int>>& gr){
    dist[tmp]=razm;
    flag[tmp]=true;
    cout<<dist[tmp]<<' '<<tmp<<' '<<razm<<"OK"<<'\n';
    for (int v:gr[tmp]){
        //cout<<v<<endl;
        if (flag[v]==false){
            razm=razm+1;
            dfs(v,razm,dist,flag,gr);
        }
    }
    //for (auto el:flag) это для вывода путей до всех вершин посещённых им
       // cout<<razm<<' ';
   // cout<<endl;
    //return;
}

int main(){
    int n,tmp,m,u,v;
    cin>>n>>m;
    vector <vector<int>> gr(n);
    for (int i=0; i<m; i++){
        cin>>u>>v;
        gr[u].push_back(v);
        gr[v].push_back(u);
    }
    int k;
    cin>>k;
    vector<vector<int>> control(k, vector<int>(n,0));
    for (int i=0; i<k; i++){
        cin>>tmp;
        vector<int>dist(n,-1);
        vector<bool>flag(n,0);
        flag[tmp]=true;
        for (int v:gr[tmp]){
            if (!flag[v]){
                int razm=1;
                dfs(v,razm,dist,flag,gr);
            }
        }
        for (int i=0; i<n; i++)
                control[tmp][i]=dist[i];
    }
    bool pokazat=true;
    for (int j=0; j<k; j++){
        for (int i=0; i<n; i++){
            cout<<control[j][i]<<' ';
        }
        cout<<'\n';
    }
    cout<<"OK";
    for (int j=0; j<n; j++){
        for (int i=0; i<k-1; i++){
            if (control[i][j]!=control[j][i+1])
                pokazat=false;
        }
        if (pokazat=true)
            cout<<j<<'\n';
    }
    if (pokazat==false)
        cout<<'-';
}


