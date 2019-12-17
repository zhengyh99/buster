<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1576569011806" ID="ID_1745680045" MODIFIED="1576569132362" TEXT="MMO&#x5ba2;&#x6237;&#x7aef;">
<node CREATED="1576569160280" ID="ID_1780568725" MODIFIED="1576569176319" POSITION="right" TEXT="Grid(&#x5750;&#x6807;&#x7f51;&#x683c;)">
<node CREATED="1576569176320" ID="ID_1277665867" MODIFIED="1576569183104" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576569193630" ID="ID_99961164" MODIFIED="1576569220812" TEXT="Gid:int&#xff08;&#x7f51;&#x683c;ID)"/>
<node CREATED="1576569234465" ID="ID_1985754933" MODIFIED="1576569236549" TEXT="Left:int&#xff08;&#x5de6;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576569246393" ID="ID_1669276866" MODIFIED="1576569248130" TEXT="Right:int  (&#x53f3;&#x8fb9;&#x754c;&#x5750;&#x6807;)"/>
<node CREATED="1576569255321" ID="ID_1419489167" MODIFIED="1576569257149" TEXT="Top:int  (&#x4e0a;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576569273037" ID="ID_999702985" MODIFIED="1576569274613" TEXT="Bottom:int  (&#x4e0b;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576569285998" ID="ID_1051299262" MODIFIED="1576569289073" TEXT="PlayerMap:map[int]bool  (&#x7f51;&#x683c;&#x4e2d;&#x7684;&#x73a9;&#x5bb6;&#x96c6;&#x5408; &#x3010;&#x5b58;&#x50a8;&#x6570;&#x636e;&#x7c7b;&#x578b;&#xff1a;map[playerID]bool&#x3011;)"/>
<node CREATED="1576569300086" ID="ID_1822874930" MODIFIED="1576569301860" TEXT="PMLock:sync.RWMutex(&#x98ce;&#x683c;&#x4e2d;&#x73a9;&#x5bb6;&#x96c6;&#x5408;CRUD&#x4fdd;&#x62a4;&#x9501;)"/>
</node>
<node CREATED="1576569183959" ID="ID_150947790" MODIFIED="1576569189387" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576569327743" ID="ID_2853239" MODIFIED="1576569341277" TEXT="NewGrid(gid,left,right,top,bottom int)*Grid(&#x65b0;&#x5efa;&#x4e00;&#x4e2a;&#x7f51;&#x683c;)"/>
<node CREATED="1576569343353" ID="ID_1381537577" MODIFIED="1576569354355" TEXT="(g *Grid)Add(playerID int32)  ( &#x5411;&#x7f51;&#x683c;&#x4e2d;&#x6dfb;&#x52a0;&#x73a9;&#x5bb6;)"/>
<node CREATED="1576569355037" ID="ID_943803148" MODIFIED="1576569364765" TEXT="(g *Grid)Remove(playerID int32) (&#x5c06;PlayerID&#x5750;&#x7f51;&#x683c;&#x73a9;&#x5bb6;&#x96c6;&#x5408;&#x4e2d;&#x5220;&#x9664;)"/>
<node CREATED="1576569365368" ID="ID_1674770689" MODIFIED="1576569375848" TEXT="(g *Grid)GetPlayerIDs()[]int32  (&#x8fd4;&#x56de;&#x7f51;&#x683c;&#x6240;&#x6709;&#x73a9;&#x5bb6;ID)"/>
</node>
</node>
<node CREATED="1576569382809" ID="ID_747669327" MODIFIED="1576569404501" POSITION="left" TEXT="AOIManager(&#x5174;&#x8da3;&#x533a;&#x57df;&#x7ba1;&#x7406;)">
<node CREATED="1576569406248" ID="ID_489375764" MODIFIED="1576569409703" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576569416660" ID="ID_931932062" MODIFIED="1576569452648" TEXT="Left:int&#xff08;&#x5de6;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576569453262" ID="ID_531735055" MODIFIED="1576569464296" TEXT="Right:int  (&#x53f3;&#x8fb9;&#x754c;&#x5750;&#x6807;)"/>
<node CREATED="1576569464923" ID="ID_542244147" MODIFIED="1576569476955" TEXT="GridsX:int  (x&#x8f74;&#x65b9;&#x5411;&#x53ef;&#x4ee5;&#x653e;&#x7f6e;&#x7f51;&#x683c;&#x7684;&#x4e2a;&#x6570; )"/>
<node CREATED="1576569477401" ID="ID_1761587334" MODIFIED="1576569496267" TEXT="Top:int  (&#x4e0a;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576569512705" ID="ID_1231749305" MODIFIED="1576569516026" TEXT="Bottom:int  (&#x4e0b;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576569521652" ID="ID_1567900287" MODIFIED="1576569524732" TEXT="GridsY:int  (y&#x8f74;&#x65b9;&#x5411;&#x53ef;&#x4ee5;&#x653e;&#x7f6e;&#x7f51;&#x683c;&#x7684;&#x4e2a;&#x6570; )"/>
<node CREATED="1576569538549" ID="ID_1578351988" MODIFIED="1576569612457" TEXT="GridMap:map[int]*Grid&#xff08;&#x6309;&#x81ea;&#x5de6;&#x81f3;&#x53f3;&#xff0c;&#x81ea;&#x4e0a;&#x81f3;&#x4e0b;&#x7684;&#x987a;&#x5e8f;&#x5b58;&#x653e;&#x533a;&#x57df;&#x7684;&#x7f51;&#x683c;&#xff09;"/>
</node>
<node CREATED="1576569410226" ID="ID_1659510609" MODIFIED="1576569413059" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576569637849" ID="ID_614628292" MODIFIED="1576569660872" TEXT="NewAOIManager(left, right, top, bottom, gridsX, gridsY int) *AOIManager &#x65b0;&#x5efa;&#x4e00;&#x4e2a;AOIManager"/>
</node>
</node>
</node>
</map>
