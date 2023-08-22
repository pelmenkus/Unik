#include<bits/stdc++.h>
using namespace std;

void dfs (int& v,int p, vector <vector<int>>& g, vector<bool>& used, vector<int>& tin, vector<int>& fup,int& timer,int& bridges) {
	used[v] = true;
	tin[v] = fup[v] = timer++;
	for (size_t i=0; i<g[v].size(); ++i) {
		int to = g[v][i];
		if (to == p)  continue;
		if (used[to])
			fup[v] = min (fup[v], tin[to]);
		else {
			dfs (to, v,g,used,tin,fup,timer,bridges);
			fup[v] = min (fup[v], fup[to]);
			if (fup[to] > tin[v])
				bridges++;
		}
	}
}

int main(){
    int n,m;
    cin>>n>>m;
    vector <vector<int>> gr(n);
    vector<bool> flag(n,false);
    for (int i=0; i<m; i++){
        int u,v;
        cin>>u>>v;
        gr[u].push_back(v);
        gr[v].push_back(u);
    }
    int bridges=0;
    vector<int>tin(n,0);
    vector<int>fup(n,0);
    for (int i=0; i<n; i++){
        if (flag[i]==false){
            int timer=0;
            int param=-1;
            dfs(i,param,gr,flag,tin,fup,timer,bridges);
        }
    }
    cout<<bridges;
}


