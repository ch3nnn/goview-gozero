var k=(t,e,n)=>new Promise((c,o)=>{var l=i=>{try{f(n.next(i))}catch(g){o(g)}},v=i=>{try{f(n.throw(i))}catch(g){o(g)}},f=i=>i.done?c(i.value):Promise.resolve(i.value).then(l,v);f((n=n.apply(t,e)).next())});import{M as H}from"./index-f81becf1.js";import{_ as G}from"./index-e7e74142.js";import{u as U,b as V,f as x}from"./chartEditStore-32c50897.js";import{u as J,a as O}from"./chartLayoutStore-a3dfb071.js";import{_ as Y,L as q,f as D,d as L,e as z,c as Q}from"./index-5e1011db.js";import{cL as W,cM as A,cN as X,cO as Z,g as j,cP as tt,cQ as et,cR as at,cS as nt,cT as st,d as ot,j as rt,a1 as ct,J as lt,ad as it,r as b,o as _,c as h,n as C,F as dt,q as _t,b as d,u as m,w as u,e as P,t as w,D as T,a6 as $,av as ut,cU as ft,aK as gt,s as pt,x as ht,h as mt}from"./index-312a31dd.js";import{a as Ct,b as vt,l as yt,g as bt}from"./plugin-aa00071a.js";import{c as E}from"./index-d277f175.js";import{i as Et}from"./icon-ce016406.js";import"./index-ffd25867.js";import"./SettingItem-dba63559.js";/* empty css                                                                      */import"./SettingItemBox-19ac3d74.js";import"./CollapseItem.vue_vue_type_script_setup_true_lang-da631111.js";import"./index.esm.min-f20ed2aa.js";import"./fileTypeEnum-21359a08.js";function St(t){var e=t==null?0:t.length;return e?t[e-1]:void 0}var kt=St;function xt(t,e,n){var c=-1,o=t.length;e<0&&(e=-e>o?0:o+e),n=n>o?o:n,n<0&&(n+=o),o=e>n?0:n-e>>>0,e>>>=0;for(var l=Array(o);++c<o;)l[c]=t[c+e];return l}var Ot=xt,Dt=W,Lt=Ot;function Pt(t,e){return e.length<2?t:Dt(t,Lt(e,0,-1))}var wt=Pt,Tt=A,$t=kt,At=wt,It=X;function Kt(t,e){return e=Tt(e,t),t=At(t,e),t==null||delete t[It($t(e))]}var Bt=Kt,Ft=Z;function Nt(t){return Ft(t)?void 0:t}var Rt=Nt,Mt=tt,Ht=st,Gt=Bt,Ut=A,Vt=et,Jt=Rt,Yt=at,qt=nt,zt=1,Qt=2,Wt=4,Xt=Yt(function(t,e){var n={};if(t==null)return n;var c=!1;e=Mt(e,function(l){return l=Ut(l,t),c||(c=l.length>1),l}),Vt(t,qt(t),n),c&&(n=Ht(n,zt|Qt|Wt,Jt));for(var o=e.length;o--;)Gt(n,e[o]);return n}),Zt=Xt;const jt=j(Zt),te=t=>(pt("data-v-fa3df045"),t=t(),ht(),t),ee={class:"go-content-charts-item-animation-patch"},ae=["onDragstart","onDragend","onDblclick","onClick"],ne={class:"list-header"},se={class:"list-center go-flex-center go-transition",draggable:"true"},oe={class:"list-bottom"},re={key:0,class:"list-model"},ce=["onClick"],le=te(()=>C("span",null,"删除",-1)),ie=ot({__name:"index",props:{menuOptions:{type:Array,default:()=>[]}},emits:["deletePhoto"],setup(t,{emit:e}){const n=U(),{TrashIcon:c}=Et.ionicons5,o=e,l=J(),v=rt(),f=a=>!a.disabled&&a.package===V.PHOTOS&&a.category===q.PRIVATE,i=ct(()=>l.getChartType),g=(a,s)=>{s.disabled||(E(s.chartKey,D(s)),E(s.conKey,L(s)),a.dataTransfer.setData(ft.DRAG_KEY,gt(jt(s,["image"]))),n.setEditCanvas(x.IS_CREATE,!0))},I=()=>{n.setEditCanvas(x.IS_CREATE,!1)},K=a=>k(this,null,function*(){if(!a.disabled)try{Ct(),E(a.chartKey,D(a)),E(a.conKey,L(a));let s=yield z(a);a.redirectComponent&&(a.dataset&&(s.option.dataset=a.dataset),s.chartConfig.title=a.title,s.chartConfig.chartFrame=a.chartFrame),n.addComponentList(s,!1,!0),n.setTargetSelectChart(s.id),vt()}catch(s){yt(),window.$message.warning("图表正在研发中, 敬请期待...")}}),B=a=>{var s;(s=a==null?void 0:a.configEvents)==null||s.addHandle(a)},F=(a,s)=>{bt({message:"是否删除此图片？",transformOrigin:"center",onPositiveCallback:()=>{const y=Q();o("deletePhoto",a,s),y.deletePhotos(a,s)}})};return lt(()=>i.value,a=>{a===O.DOUBLE&&it(()=>{v.value.classList.add("miniAnimation")})}),(a,s)=>{const y=b("n-ellipsis"),S=b("n-text"),N=b("n-icon"),R=b("n-button");return _(),h("div",ee,[C("div",{ref_key:"contentChartsItemBoxRef",ref:v,class:ut(["go-content-charts-item-box",[i.value===m(O).DOUBLE?"double":"single"]])},[(_(!0),h(dt,null,_t(t.menuOptions,(r,M)=>(_(),h("div",{class:"item-box",key:r.title,draggable:"",onDragstart:p=>!r.disabled&&g(p,r),onDragend:p=>!r.disabled&&I,onDblclick:p=>K(r),onClick:p=>B(r)},[C("div",ne,[d(m(H),{class:"list-header-control-btn",mini:!0,disabled:!0}),d(S,{class:"list-header-text",depth:"3"},{default:u(()=>[d(y,null,{default:u(()=>[P(w(r.title),1)]),_:2},1024)]),_:2},1024)]),C("div",se,[r.icon?(_(),T(m(Y),{key:0,class:"list-img",icon:r.icon,color:"#999",width:"48",style:{height:"auto"}},null,8,["icon"])):(_(),T(m(G),{key:1,class:"list-img",chartConfig:r},null,8,["chartConfig"]))]),C("div",oe,[d(S,{class:"list-bottom-text",depth:"3"},{default:u(()=>[d(y,{style:{"max-width":"90%"}},{default:u(()=>[P(w(r.title),1)]),_:2},1024)]),_:2},1024)]),r.disabled?(_(),h("div",re)):$("",!0),f(r)?(_(),h("div",{key:1,class:"list-tools go-transition",onClick:p=>F(r,M)},[d(R,{text:"",type:"default",color:"#ffffff"},{icon:u(()=>[d(N,null,{default:u(()=>[d(m(c))]),_:1})]),default:u(()=>[le]),_:1})],8,ce)):$("",!0)],40,ae))),128))],2)])}}});const De=mt(ie,[["__scopeId","data-v-fa3df045"]]);export{De as default};
