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
<node CREATED="1576679599675" ID="ID_323942761" MODIFIED="1576679622651" TEXT="func (dp *DataPack) UnPack(data []byte) (iface.IMessage, error)   &#x62c6;&#x5305;"/>
</node>
</node>
</map>
