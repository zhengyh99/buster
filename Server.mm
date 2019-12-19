<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1576656004368" ID="ID_862665061" MODIFIED="1576678813418" TEXT="TCP Server">
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
</node>
</map>
