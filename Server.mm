<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1576656004368" ID="ID_862665061" MODIFIED="1576819746329" TEXT="TCP Server">
<node CREATED="1576676258712" ID="ID_228078391" MODIFIED="1576676685746" POSITION="right" TEXT="Message(&#x6d88;&#x606f;&#xff09;">
<font NAME="SansSerif" SIZE="14"/>
<node CREATED="1576676723979" ID="ID_704466269" MODIFIED="1576676727932" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576676739537" ID="ID_212281363" MODIFIED="1576676767299" TEXT="ID:int32 &#x6d88;&#x606f;ID"/>
<node CREATED="1576676768427" ID="ID_537167105" MODIFIED="1576676796886" TEXT="DataLen:int32 &#x6d88;&#x606f;&#x957f;&#x5ea6;"/>
<node CREATED="1576676814364" ID="ID_627916687" MODIFIED="1576676828462" TEXT="Data:[]byte &#x6d88;&#x606f;&#x5185;&#x5bb9;"/>
</node>
<node CREATED="1576676728727" ID="ID_1726633905" MODIFIED="1576676731929" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576676846958" ID="ID_1268894281" MODIFIED="1576676878176" TEXT="NewMessage(id uint32, data []byte) *Message &#x65b0;&#x5efa;&#x4e00;&#x4e2a;&#x6d88;&#x606f;"/>
<node CREATED="1576676949646" ID="ID_196974653" MODIFIED="1576676978862" TEXT="GET &amp; SET &#x7ed3;&#x6784;&#x4f53;&#x5404;&#x5c5e;&#x6027;&#x7684;&#x65b9;&#x6cd5;"/>
</node>
</node>
<node CREATED="1576678818079" ID="ID_548744966" MODIFIED="1576679638234" POSITION="left" TEXT="DataPack&#xff08;&#x6d88;&#x606f;&#x5305;&#xff09;">
<node CREATED="1576679568227" ID="ID_1147529457" MODIFIED="1576679578982" TEXT="(dp *DataPack) GetHeadLen() uint32  &#x83b7;&#x53d6;&#x5305;&#x5934;&#x7684;&#x957f;&#x5ea6;"/>
<node CREATED="1576679593020" ID="ID_1975712061" MODIFIED="1576679598499" TEXT="func (dp *DataPack) Pack(msg iface.IMessage) (data []byte, err error)   &#x5c01;&#x5305;"/>
<node CREATED="1576679599675" ID="ID_323942761" MODIFIED="1576761668549" TEXT="func (dp *DataPack) UnPack(conn net.Conn) (iface.IMessage, error)   &#x62c6;&#x5305;"/>
</node>
<node CREATED="1576761615424" ID="ID_417456807" MODIFIED="1576762322578" POSITION="right" TEXT="Request &#xff08;&#x6d88;&#x606f;&#x8bf7;&#x6c42;&#xff09;">
<node CREATED="1576762322578" ID="ID_1050118717" MODIFIED="1576762327026" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576762581548" ID="ID_1445398140" MODIFIED="1576762611112" TEXT="conn:iface.IConnection &#x5c01;&#x88c5;&#x8fde;&#x63a5;"/>
<node CREATED="1576762611933" ID="ID_696981706" MODIFIED="1576762644813" TEXT="msg:iface.IMessage &#x5c01;&#x88c5;&#x6d88;&#x606f;"/>
</node>
<node CREATED="1576762327784" ID="ID_467751941" MODIFIED="1576762331677" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576762749738" ID="ID_74609390" MODIFIED="1576762764110" TEXT="(r *Request) GetConnection() iface.IConnection  &#x83b7;&#x53d6;&#x8fde;&#x63a5;&#x5bf9;&#x8c61;"/>
<node CREATED="1576762764922" ID="ID_29282615" MODIFIED="1576762789780" TEXT="(r *Request) GetData() []byte   &#x83b7;&#x53d6;&#x6d88;&#x606f;&#x4f53;&#x6570;&#x636e;"/>
<node CREATED="1576762806430" ID="ID_379859489" MODIFIED="1576762816742" TEXT="(r *Request) GetMsgID() uint32 &#x83b7;&#x53d6;&#x6d88;&#x606f;ID"/>
<node CREATED="1576762817951" ID="ID_794408028" MODIFIED="1576763988031" TEXT="(r *Request) Send(msgID uint32, data []byte) error  &#x5c06;&#x6d88;&#x606f; &#x53d1;&#x9001;&#x7ed9;&#x5ba2;&#x6237;&#x7aef;">
<icon BUILTIN="button_cancel"/>
</node>
</node>
</node>
<node CREATED="1576763972252" ID="ID_973722362" MODIFIED="1576764255276" POSITION="left" TEXT="ConnManager&#xff08;&#x94fe;&#x63a5;&#x7ba1;&#x7406;&#xff09;">
<node CREATED="1576764255276" ID="ID_723918424" MODIFIED="1576764272032" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576764293064" ID="ID_1345066756" MODIFIED="1576764331711" TEXT="connections:map[uint32]iface.IConnection &#x6240;&#x6709;&#x53ef;&#x7ba1;&#x7406;&#x94fe;&#x63a5;&#x7684;&#x96c6;&#x5408;"/>
<node CREATED="1576764346910" ID="ID_122752224" MODIFIED="1576764394200" TEXT="connLock : sync.RWMutex &#x94fe;&#x63a5;&#x96c6;&#x5408;&#x7684;CRUD&#x64cd;&#x4f5c;&#x4e50; &#x9501;"/>
</node>
<node CREATED="1576764272976" ID="ID_590854871" MODIFIED="1576764275198" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576764440479" ID="ID_1539941043" MODIFIED="1576764459483" TEXT="NewConnManager() iface.IConnManager   &#x65b0;&#x5efa;&#x4e00;&#x4e2a;&#x94fe;&#x63a5;&#x7ba1;&#x7406;"/>
<node CREATED="1576764462174" ID="ID_205849405" MODIFIED="1576764508643" TEXT="(cm *ConnManager) Add(conn iface.IConnection)  &#x6dfb;&#x52a0;&#x94fe;&#x63a5;"/>
<node CREATED="1576764491982" ID="ID_629589403" MODIFIED="1576764523858" TEXT=" (cm *ConnManager) Remove(conn iface.IConnection)  &#x5220;&#x9664;&#x94fe;&#x63a5;"/>
<node CREATED="1576764555380" ID="ID_509580557" MODIFIED="1576764569712" TEXT="(cm *ConnManager) Get(connID uint32) (iface.IConnection, error)  &#x901a;&#x8fc7;&#x94fe;&#x63a5;ID&#x8fd4;&#x56de;&#x94fe;&#x63a5;&#x5bf9;&#x8c61;"/>
<node CREATED="1576764584955" ID="ID_1188809977" MODIFIED="1576764598086" TEXT="(cm *ConnManager) Num() int  &#x8fd4;&#x56de;&#x53ef;&#x7ba1;&#x7406;&#x94fe;&#x63a5;&#x6570;&#x91cf;"/>
<node CREATED="1576764601021" ID="ID_598089463" MODIFIED="1576764623001" TEXT=" (cm *ConnManager) ClearAll()   &#x6e05;&#x9664;&#x6240;&#x6709;&#x94fe;&#x63a5;"/>
</node>
</node>
<node CREATED="1576764673418" ID="ID_128517967" MODIFIED="1576764702574" POSITION="right" TEXT="BaseRouter &#xff08;&#x57fa;&#x7840;&#x8def;&#x7531;&#xff09;">
<node CREATED="1576764735559" ID="ID_1787984101" MODIFIED="1576764853056" TEXT="(br *BaseRouter) PreHandle(r iface.IRequest)  &#x8def;&#x7531;&#x4e1a;&#x52a1;&#x9884;&#x5904;&#x7406;"/>
<node CREATED="1576764764160" ID="ID_213730536" MODIFIED="1576764798829" TEXT="(br *BaseRouter) Handle(r iface.IRequest) &#x8def;&#x7531;&#x4e1a;&#x52a1;&#x5904;&#x7406;  "/>
<node CREATED="1576764799257" ID="ID_225774680" MODIFIED="1576764826712" TEXT="(br *BaseRouter) PostHandle(r iface.IRequest )  &#x8def;&#x7531;&#x4e1a;&#x52a1;&#x540e;&#x7eed;&#x5904;&#x7406; "/>
</node>
<node CREATED="1576819747486" ID="ID_84068070" MODIFIED="1576820104128" POSITION="left" TEXT="MessageHandler&#xff08;&#x6d88;&#x606f;&#x5904;&#x7406;&#xff09;">
<node CREATED="1576819870361" ID="ID_418242748" MODIFIED="1576819872804" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576819879696" ID="ID_1756330284" MODIFIED="1576819994394" TEXT="MsgRouters:map[uint32]iface.IRouter &#x6d88;&#x606f;ID&#x548c;&#x6d88;&#x606f;&#x5904;&#x7406;&#x8def;&#x7531;&#x7684;&#x5bf9;&#x5e94;&#x96c6;&#x5408;"/>
<node CREATED="1576819971333" ID="ID_606740124" MODIFIED="1576820001545" TEXT="TaskQueue:[]chan iface.IRequest &#x4efb;&#x52a1;&#x961f;&#x5217; "/>
<node CREATED="1576820012630" ID="ID_670273675" MODIFIED="1576820039364" TEXT="TaskPullSize:uint32 &#x4efb;&#x52a1;&#x6c60;&#x4e2d;&#x4efb;&#x52a1;&#x6570;&#x91cf; "/>
</node>
<node CREATED="1576819873470" ID="ID_37117203" MODIFIED="1576819876210" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576820077383" ID="ID_1824134524" MODIFIED="1576820097038" TEXT="NewMsgHandler() *MsgHandler &#x65b0;&#x5efa;&#x4e00;&#x4e2a;&#x6d88;&#x606f;&#x5904;&#x7406;"/>
<node CREATED="1576820354608" ID="ID_730454256" MODIFIED="1576820365696" TEXT="&#x8def;&#x7531;&#x7ba1;&#x7406;&#x65b9;&#x6cd5;">
<node CREATED="1576820175673" ID="ID_289260619" MODIFIED="1576820231839" TEXT="(mh *MsgHandler) DoMsgHandler(request iface.IRequest) error  &#x6267;&#x884c;request&#x4e2d;&#x7684;msgID&#x5bf9;&#x5e94;&#x7684;&#x6d88;&#x606f;&#x5904;&#x7406;&#x8def;&#x7531;"/>
<node CREATED="1576820263853" ID="ID_713322341" MODIFIED="1576820302364" TEXT="(mh *MsgHandler) AddRouter(msgID uint32, router iface.IRouter) error  &#x6dfb;&#x52a0;&#x8def;&#x7531; &#x5230;MsgRouters &#x5c5e;&#x6027;&#x4e2d;"/>
</node>
<node CREATED="1576820367276" ID="ID_956553681" MODIFIED="1576820372363" TEXT="&#x4efb;&#x52a1;&#x6c60;&#x7ba1;&#x7406;&#x65b9;&#x6cd5;">
<node CREATED="1576820400485" ID="ID_932330327" MODIFIED="1576820413484" TEXT="(mh *MsgHandler) OpenTaskPool()  &#x6253;&#x5f00;&#x8def;&#x7531;&#x6c60;"/>
<node CREATED="1576820415337" ID="ID_95810106" MODIFIED="1576820482304" TEXT="(mh *MsgHandler) RunTask(taskID uint32, taskQueue chan iface.IRequest) &#x5355;&#x4e2a;&#x4efb;&#x52a1;&#x5904;&#x7406;"/>
<node CREATED="1576820525681" ID="ID_296581676" MODIFIED="1576820540627" TEXT="(mh *MsgHandler) SendToTask(request iface.IRequest)  &#x4efb;&#x52a1;&#x5206;&#x914d;"/>
</node>
</node>
</node>
<node CREATED="1576820810284" ID="ID_1764724722" MODIFIED="1576820826218" POSITION="right" TEXT="Connection(&#x7f51;&#x8def;&#x94fe;&#x63a5;&#xff09;">
<node CREATED="1576820832390" ID="ID_1421476370" MODIFIED="1576820835843" TEXT="&#x5c5e;&#x6027;">
<node CREATED="1576820851858" ID="ID_673403097" MODIFIED="1576820882323" TEXT="MyServer:iface.IServer &#x7f51;&#x7edc;&#x94fe;&#x63a5;&#x96b6;&#x5c5e;&#x7684;&#x670d;&#x52a1;&#x5668;"/>
<node CREATED="1576820883157" ID="ID_264186440" MODIFIED="1576821077617" TEXT="Conn:*net.TCPConn &#x94fe;&#x63a5;Socket TCP &#x5957;&#x63a5;&#x5b57;"/>
<node CREATED="1576821034519" ID="ID_333154406" MODIFIED="1576821072145" TEXT="ConnID:uint32 &#x94fe;&#x63a5;ID"/>
<node CREATED="1576821041745" ID="ID_1988029517" MODIFIED="1576821065685" TEXT="isClose:bool &#x94fe;&#x63a5;&#x72b6;&#x6001;"/>
<node CREATED="1576821107976" ID="ID_1043205621" MODIFIED="1576821266992" TEXT="MsgChan:chan []byte &#x8bfb;&#x5199;&#x64cd;&#x4f5c;&#x6d88;&#x606f;&#x901a;&#x9053;"/>
<node CREATED="1576821244786" ID="ID_1612917376" MODIFIED="1576821307750" TEXT="ExitChan:chan bool &#x94fe;&#x63a5;&#x9000;&#x51fa;/&#x505c;&#x6b62;&#x6d88;&#x606f;&#x901a;&#x9053;"/>
<node CREATED="1576821328777" ID="ID_1223065976" MODIFIED="1576821346685" TEXT="MsgHandler:iface.IMsgHandler &#x6d88;&#x606f;&#x5904;&#x7406;"/>
<node CREATED="1576821361765" ID="ID_1272603506" MODIFIED="1576821377398" TEXT="property map[string]interface{} &#x94fe;&#x63a5;&#x7684;&#x81ea;&#x5b9a;&#x5c5e;&#x6027;&#x96c6;&#x5408;"/>
<node CREATED="1576821378414" ID="ID_700967606" MODIFIED="1576821414310" TEXT="propertyLock sync.RWMutex &#x94fe;&#x63a5;&#x81ea;&#x5b9a;&#x4e49;&#x5c5e;&#x6027;CRUD&#x64cd;&#x4f5c;&#x9501;"/>
</node>
<node CREATED="1576820836285" ID="ID_1542417450" MODIFIED="1576824947934" TEXT="&#x65b9;&#x6cd5;">
<node CREATED="1576825196037" ID="ID_1889492627" MODIFIED="1576825226757" TEXT="(c *Connection) Start()  &#x6253;&#x5f00;&#x94fe;&#x63a5;">
<node CREATED="1576825279592" ID="ID_437627926" MODIFIED="1576825344990" TEXT="(c *Connection) StartRead()  &#x8bfb;&#x53d6;&#x5ba2;&#x6237;&#x7aef;&#x6570;&#x636e; ">
<node CREATED="1576825597562" ID="ID_1730427709" MODIFIED="1576825671311" TEXT="(c *Connection) getMsg() (iface.IMessage, error)  &#x62c6;&#x5305;&#x3001;&#x89e3;&#x6790;&#x5ba2;&#x6237;&#x7aef;&#x6570;&#x636e;"/>
</node>
<node CREATED="1576825306087" ID="ID_1390838285" MODIFIED="1576825340117" TEXT="(c *Connection) StartWrite()  &#x5411;&#x5ba2;&#x6237;&#x7aef;&#x5199;&#x5165;&#x6570;&#x636e; "/>
</node>
<node CREATED="1576825227772" ID="ID_88832460" MODIFIED="1576825242570" TEXT="(c *Connection) Stop()  &#x5173;&#x95ed;&#x94fe;&#x63a5;"/>
<node CREATED="1576825495573" ID="ID_1188605337" MODIFIED="1576825558354" TEXT="(c *Connection) Send(msgID uint32, msgData []byte) error  &#x5411;&#x5ba2;&#x6237;&#x7aef;&#x53d1;&#x9001;Message"/>
<node CREATED="1576824738285" ID="ID_671598174" MODIFIED="1576824749887" TEXT="GET &#x65b9;&#x6cd5;">
<node CREATED="1576824751214" ID="ID_1851911060" MODIFIED="1576824803441" TEXT="(c *Connection) GetTCPConnection() *net.TCPConn  &#x83b7;&#x53d6; &#x539f;&#x751f;SocketTCP&#x5957;&#x63a5;&#x5b57;"/>
<node CREATED="1576824821600" ID="ID_693585701" MODIFIED="1576824832917" TEXT="(c *Connection) GetConnID() uint32 &#x83b7;&#x53d6;&#x94fe;&#x63a5;ID"/>
<node CREATED="1576824853557" ID="ID_1978100775" MODIFIED="1576824881308" TEXT="(c *Connection) RemoteAddr() net.Addr &#x83b7;&#x53d6;&#x5ba2;&#x6237;&#x7aef; IP&#x5730;&#x5740;&#x548c;&#x7aef;&#x53e3;&#x53f7;"/>
</node>
<node CREATED="1576824947935" ID="ID_1687675544" MODIFIED="1576824961418" TEXT="&#x81ea;&#x5b9a;&#x4e49;&#x5c5e;&#x6027;&#x64cd;&#x4f5c;">
<node CREATED="1576824963322" ID="ID_1360105991" MODIFIED="1576824986511" TEXT="(c *Connection) SetProperty(key string, value interface{})   &#x8bbe;&#x7f6e;&#x5c5e;&#x6027;"/>
<node CREATED="1576825006987" ID="ID_1689405184" MODIFIED="1576825015079" TEXT="(c *Connection) GetProperty(key string) interface{}  &#x83b7;&#x53d6;&#x5c5e;&#x6027;&#x503c;"/>
<node CREATED="1576825033365" ID="ID_271499444" MODIFIED="1576825039382" TEXT="(c *Connection) RemoveProperty(key string)  &#x5220;&#x9664;&#x5c5e;&#x6027;"/>
</node>
</node>
</node>
</node>
</map>
