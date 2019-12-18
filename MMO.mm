<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1576569011806" ID="ID_1745680045" MODIFIED="1576655984504" TEXT="MMO&#x670d;&#x52a1;&#x5668;">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576569160280" ID="ID_1780568725" MODIFIED="1576643673771" POSITION="right" TEXT="Grid(&#x5750;&#x6807;&#x7f51;&#x683c;)">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576569176320" ID="ID_1277665867" MODIFIED="1576643673771" TEXT="&#x5c5e;&#x6027;">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576569193630" ID="ID_99961164" MODIFIED="1576643673771" TEXT="Gid:int&#xff08;&#x7f51;&#x683c;ID)">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569234465" ID="ID_1985754933" MODIFIED="1576643673771" TEXT="Left:int&#xff08;&#x5de6;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569246393" ID="ID_1669276866" MODIFIED="1576643673771" TEXT="Right:int  (&#x53f3;&#x8fb9;&#x754c;&#x5750;&#x6807;)">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569255321" ID="ID_1419489167" MODIFIED="1576643673771" TEXT="Top:int  (&#x4e0a;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569273037" ID="ID_999702985" MODIFIED="1576643673771" TEXT="Bottom:int  (&#x4e0b;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569285998" ID="ID_1051299262" MODIFIED="1576643673771" TEXT="PlayerMap:map[int]bool  (&#x7f51;&#x683c;&#x4e2d;&#x7684;&#x73a9;&#x5bb6;&#x96c6;&#x5408; &#x3010;&#x5b58;&#x50a8;&#x6570;&#x636e;&#x7c7b;&#x578b;&#xff1a;map[playerID]bool&#x3011;)">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569300086" ID="ID_1822874930" MODIFIED="1576643673771" TEXT="PMLock:sync.RWMutex(&#x98ce;&#x683c;&#x4e2d;&#x73a9;&#x5bb6;&#x96c6;&#x5408;CRUD&#x4fdd;&#x62a4;&#x9501;)">
<font NAME="SansSerif" SIZE="14"/>
</node>
</node>
<node CREATED="1576569183959" ID="ID_150947790" MODIFIED="1576643673771" TEXT="&#x65b9;&#x6cd5;">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576569327743" ID="ID_2853239" MODIFIED="1576643673771" TEXT="NewGrid(gid,left,right,top,bottom int)*Grid(&#x65b0;&#x5efa;&#x4e00;&#x4e2a;&#x7f51;&#x683c;)">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569343353" ID="ID_1381537577" MODIFIED="1576643673771" TEXT="(g *Grid)Add(playerID int32)  ( &#x5411;&#x7f51;&#x683c;&#x4e2d;&#x6dfb;&#x52a0;&#x73a9;&#x5bb6;)">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569355037" ID="ID_943803148" MODIFIED="1576643673771" TEXT="(g *Grid)Remove(playerID int32) (&#x5c06;PlayerID&#x5750;&#x7f51;&#x683c;&#x73a9;&#x5bb6;&#x96c6;&#x5408;&#x4e2d;&#x5220;&#x9664;)">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569365368" ID="ID_1674770689" MODIFIED="1576643673771" TEXT="(g *Grid)GetPlayerIDs()[]int32  (&#x8fd4;&#x56de;&#x7f51;&#x683c;&#x6240;&#x6709;&#x73a9;&#x5bb6;ID)">
<font NAME="SansSerif" SIZE="14"/>
</node>
</node>
</node>
<node CREATED="1576569382809" ID="ID_747669327" MODIFIED="1576643673771" POSITION="left" TEXT="AOIManager(&#x5174;&#x8da3;&#x533a;&#x57df;&#x7ba1;&#x7406;)">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576569406248" ID="ID_489375764" MODIFIED="1576643673771" TEXT="&#x5c5e;&#x6027;">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576569416660" ID="ID_931932062" MODIFIED="1576643673770" TEXT="Left:int&#xff08;&#x5de6;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569453262" ID="ID_531735055" MODIFIED="1576643673770" TEXT="Right:int  (&#x53f3;&#x8fb9;&#x754c;&#x5750;&#x6807;)">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569464923" ID="ID_542244147" MODIFIED="1576643673770" TEXT="GridsX:int  (x&#x8f74;&#x65b9;&#x5411;&#x53ef;&#x4ee5;&#x653e;&#x7f6e;&#x7f51;&#x683c;&#x7684;&#x4e2a;&#x6570; )">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569477401" ID="ID_1761587334" MODIFIED="1576643673770" TEXT="Top:int  (&#x4e0a;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569512705" ID="ID_1231749305" MODIFIED="1576643673770" TEXT="Bottom:int  (&#x4e0b;&#x8fb9;&#x754c;&#x5750;&#x6807;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569521652" ID="ID_1567900287" MODIFIED="1576643673770" TEXT="GridsY:int  (y&#x8f74;&#x65b9;&#x5411;&#x53ef;&#x4ee5;&#x653e;&#x7f6e;&#x7f51;&#x683c;&#x7684;&#x4e2a;&#x6570; )">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576569538549" ID="ID_1578351988" MODIFIED="1576643673770" TEXT="GridMap:map[int]*Grid&#xff08;&#x6309;&#x81ea;&#x5de6;&#x81f3;&#x53f3;&#xff0c;&#x81ea;&#x4e0a;&#x81f3;&#x4e0b;&#x7684;&#x987a;&#x5e8f;&#x5b58;&#x653e;&#x533a;&#x57df;&#x7684;&#x7f51;&#x683c;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
</node>
</node>
<node CREATED="1576569410226" ID="ID_1659510609" MODIFIED="1576643673770" TEXT="&#x65b9;&#x6cd5;">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576569637849" ID="ID_614628292" MODIFIED="1576643673769" TEXT="NewAOIManager(left, right, top, bottom, gridsX, gridsY int) *AOIManager &#x65b0;&#x5efa;&#x4e00;&#x4e2a;AOIManager">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576643220668" ID="ID_1589614202" MODIFIED="1576643283720" TEXT="Grid &#x7f51;&#x683c;  &#x76f8;&#x5173;&#x65b9;&#x6cd5;">
<node CREATED="1576642030630" ID="ID_534809872" MODIFIED="1576643673769" STYLE="fork" TEXT="(am *AOIManager) gridHeight() int &#x83b7;&#x53d6;&#x5355;&#x4e2a;Grid&#x7684;&#x9ad8;&#x5ea6;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576641992219" ID="ID_675063227" MODIFIED="1576643673769" TEXT="(am *AOIManager) gridWidth() int &#x83b7;&#x53d6;&#x5355;&#x4e2a;Grid&#x7684;&#x5bbd;&#x5ea6; ">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576642178607" ID="ID_275874411" MODIFIED="1576643673769" TEXT="(am *AOIManager) GetSurroundingsByGid(gid int) (grids []*Grid)  &#x6839;&#x636e; &#x6307;&#x5b9a;&#x7f51;&#x683c;&#x7684;gid &#x8fd4;&#x5468;&#x8fb9;&#x4e5d;&#x5bab;&#x683c;&#x96c6;&#x5408;">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576642539955" ID="ID_1712878775" MODIFIED="1576643673768" TEXT="(am *AOIManager) GetGidByPosition(x, y float32) (gid int)  &#x8fd4;&#x56de;&#x6307;&#x5b9a;&#x5750;&#x6807;&#x6240;&#x5728;Grid&#x7684;ID">
<font NAME="SansSerif" SIZE="14"/>
</node>
</node>
<node CREATED="1576643305471" ID="ID_1725206028" MODIFIED="1576643321348" TEXT="Player &#x73a9;&#x5bb6; &#x76f8;&#x5173;&#x65b9;&#x6cd5;">
<node CREATED="1576643340607" ID="ID_388720498" MODIFIED="1576643673768" STYLE="fork" TEXT="(am *AOIManager) AddPlayerToGrid(pid int32, gid int) &#x5411;Grid : gid &#x4e2d;&#x6dfb;&#x52a0;Player : pid">
<font NAME="SansSerif" SIZE="14"/>
</node>
<node CREATED="1576643477198" ID="ID_1513252647" MODIFIED="1576643694236" TEXT="(am *AOIManager) RemovePlayerFromGrid(pid int32, gid int)   &#x5c06;Grid : gid &#x4e2d;Player : pid &#x5220;&#x9664;"/>
<node CREATED="1576643644550" ID="ID_1256396645" MODIFIED="1576643899210" TEXT="(am *AOIManager) GetPlayerIDsFromGrid(gid int) (playerIDs []int32)  &#x83b7;&#x53d6;Grid : gid &#x4e2d;&#x7684; Player ID &#x96c6;&#x5408;"/>
<node CREATED="1576643968062" ID="ID_1035263236" MODIFIED="1576643987104" TEXT="(am *AOIManager) AddPlayerToPosition(pid int32, x, y float32) &#x5411;&#x5750;&#x6807;&#x7cfb; &#xff08;x,y&#xff09;&#x5bf9;&#x5e94;&#x7684;Grid:gid&#x4e2d;&#x52a0;&#x5165;Player:pid"/>
<node CREATED="1576644082102" ID="ID_1795649659" MODIFIED="1576644126110" TEXT="(am *AOIManager) RemovePlayerFromPosition(pid int32, x, y float32) &#x5c06;Player : pid &#x5750;&#x6307;&#x5b9a;&#x7684;&#x5750;&#x6807;&#xff08;x,y&#xff09; &#x4e2d;&#x5220;&#x9664;"/>
<node CREATED="1576643014552" ID="ID_1883774849" MODIFIED="1576643673766" TEXT="(am *AOIManager) GetPlayersByPosition(x, y float32) (playerIDs []int32)  &#x8fd4;&#x56de;&#x6307;&#x5b9a;&#x5750;&#x6807;&#x6240;&#x5728;Grid&#x7684;&#x5468;&#x8fb9;&#x4e5d;&#x5bab;&#x683c;&#x5185;&#x6240;&#x6709;&#x73a9;&#x5bb6;&#x7684;ID&#x96c6;&#x5408;">
<font NAME="SansSerif" SIZE="14"/>
</node>
</node>
</node>
</node>
<node CREATED="1576654515081" ID="ID_782617834" MODIFIED="1576654532449" POSITION="right" TEXT="Player(&#x73a9;&#x5bb6;)">
<node CREATED="1576654593353" ID="ID_380790216" MODIFIED="1576654595966" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576654634710" ID="ID_1470818108" MODIFIED="1576654755668" TEXT="Pid:int32 &#xff08;&#x73a9;&#x5bb6;id&#xff09;"/>
<node CREATED="1576654678649" ID="ID_919942287" MODIFIED="1576654770803" TEXT="Conn:iface.IConnection (&#x4e0e;&#x5ba2;&#x6237;&#x7aef;&#x94fe;&#x63a5;)"/>
<node CREATED="1576654781185" ID="ID_826455531" MODIFIED="1576654921958" TEXT="X:float32 &#xff08;&#x73a9;&#x5bb6;x&#x8f74;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576654852386" ID="ID_1417171822" MODIFIED="1576654929576" TEXT="Y:float32 &#xff08;&#x73a9;&#x5bb6;&#x5782;&#x76f4;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576654870572" ID="ID_907032069" MODIFIED="1576654936713" TEXT="Z:float32 &#xff08;&#x73a9;&#x5bb6;y&#x8f74;&#x5750;&#x6807;&#xff09;"/>
<node CREATED="1576654881186" ID="ID_388855486" MODIFIED="1576654948507" TEXT="V:float32 &#xff08;&#x73a9;&#x5bb6;&#x503e;&#x659c;&#x89d2;&#x5ea6;&#xff09;"/>
</node>
<node CREATED="1576654597046" ID="ID_672717644" MODIFIED="1576654600336" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576654964699" ID="ID_305114631" MODIFIED="1576654966831" TEXT="NewPlayer(conn iface.IConnection) *Player "/>
<node CREATED="1576654988466" ID="ID_1267094838" MODIFIED="1576654991832" TEXT="(p *Player) SendMsg(msgID uint32, data proto.Message) error"/>
<node CREATED="1576655240566" ID="ID_246952208" MODIFIED="1576655263857" TEXT="(p *Player) getSurroundingPlayers() []*Player  &#x83b7;&#x53d6;&#x5468;&#x56f4;&#x73a9;&#x5bb6;&#x96c6;&#x5408;"/>
<node CREATED="1576655017695" ID="ID_762683138" MODIFIED="1576655054903" TEXT="protobuf &#x4e0e;&#x5ba2;&#x6237;&#x7aef;&#x540c;&#x6b65;&#x65b9;&#x6cd5;">
<node CREATED="1576655056813" ID="ID_1171549962" MODIFIED="1576655095281" TEXT="(p *Player) SyncPid()  &#x544a;&#x77e5;&#x73a9;&#x5bb6; Pid ,&#x540c;&#x6b65;&#x670d;&#x52a1;&#x5668;&#x7aef;&#x7684; pid &#x7ed9;&#x5ba2;&#x6237;&#x7aef;"/>
<node CREATED="1576655098494" ID="ID_1801602897" MODIFIED="1576655142301" TEXT="(p *Player) BroadCastStartPosition()  &#x540c;&#x6b65;&#x73a9;&#x5bb6;&#x4f4d;&#x7f6e;&#x4fe1;&#x606f;"/>
<node CREATED="1576655177671" ID="ID_682833175" MODIFIED="1576655207239" TEXT="(p *Player) Talk(content string)  &#x63a5;&#x6536;&#x5ba2;&#x6237;&#x7aef;&#x804a;&#x5929;&#x4fe1;&#x606f;&#xff0c;&#x5e7f;&#x64ad;&#x7ed9;&#x5176;&#x5b83;&#x5728;&#x7ebf;&#x73a9;&#x5bb6;"/>
<node CREATED="1576655210726" ID="ID_985251300" MODIFIED="1576655323174" TEXT="(p *Player) SyncSurrounding()  &#x540c;&#x6b65;&#x73a9;&#x5bb6;&#x81ea;&#x5df1;&#x548c;&#x5468;&#x8fb9;&#x73a9;&#x5bb6;&#x4f4d;&#x7f6e;&#x4fe1;&#x606f;"/>
<node CREATED="1576655338718" ID="ID_1227783412" MODIFIED="1576655357384" TEXT="(p *Player) UpdatePosition(x, y, z, v float32)   &#x66f4;&#x65b0;&#x73a9;&#x5bb6;&#x4f4d;&#x7f6e;&#x4fe1;&#x606f;"/>
<node CREATED="1576655359372" ID="ID_1554722245" MODIFIED="1576655379874" TEXT="(p *Player) OffLine() &#x5e7f;&#x64ad;&#x73a9;&#x5bb6;&#x9000;&#x51fa;&#x6d88;&#x606f;"/>
</node>
</node>
</node>
<node CREATED="1576654533848" ID="ID_1177736339" MODIFIED="1576654587727" POSITION="left" TEXT="WorldManager(&#x4e16;&#x754c;&#x7ba1;&#x7406;&#xff09;">
<node CREATED="1576655599839" ID="ID_1555361276" MODIFIED="1576655602113" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576655400967" ID="ID_942497976" MODIFIED="1576655448631" TEXT="AoiMng:*AOIManager (&#x5174;&#x8da3;&#x533a;&#x57df;&#x7ba1;&#x7406;)"/>
<node CREATED="1576655450936" ID="ID_16227571" MODIFIED="1576655517629" TEXT="Players:map[int32]*Player (&#x73a9;&#x5bb6;&#x96c6;&#x5408; map[pid]*Player&#xff09;"/>
<node CREATED="1576655631851" ID="ID_345749332" MODIFIED="1576655663739" TEXT="PLock:sync.RWMutex (&#x73a9;&#x5bb6;&#x96c6;&#x5408;CRUD&#x64cd;&#x4f5c;&#x9501;)"/>
</node>
<node CREATED="1576655587944" ID="ID_250564753" MODIFIED="1576655594624" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576655561747" ID="ID_530191782" MODIFIED="1576655749848" TEXT=" (wm *WorldManager) AddPlayer(player *Player)  &#x5411;&#x5c5e;&#x6027;AoiMng&#x548c;&#x73a9;&#x5bb6;&#x96c6;&#x5408;&#x4e2d;&#x52a0;&#x5165;Player"/>
<node CREATED="1576655774720" ID="ID_403869331" MODIFIED="1576655810595" TEXT="func (wm *WorldManager) RemovePlayer(pid int32) &#x6839;&#x636e; pid &#x79fb;&#x9664;Player"/>
<node CREATED="1576655844395" ID="ID_1030205946" MODIFIED="1576655871096" TEXT="(wm *WorldManager) GetPlayer(pid int32) *Player &#x6839;&#x636e;pid &#x8fd4;&#x56de;Player"/>
<node CREATED="1576655883909" ID="ID_1600514675" MODIFIED="1576655924978" TEXT="(wm *WorldManager) GetAllPlayer() (players []*Player)    &#x56de;&#x73a9;&#x5bb6;&#x4e16;&#x754c;&#x4e2d;&#x6240;&#x6709;Player"/>
</node>
</node>
</node>
</map>
