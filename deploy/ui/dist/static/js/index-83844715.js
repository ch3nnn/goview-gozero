var O=(e,s,u)=>new Promise((m,h)=>{var D=t=>{try{g(u.next(t))}catch(_){h(_)}},d=t=>{try{g(u.throw(t))}catch(_){h(_)}},g=t=>t.done?m(t.value):Promise.resolve(t.value).then(D,d);g((u=u.apply(e,s)).next())});import{d as T,l as V,H,Y as S,j,r as c,o as r,c as C,b as a,w as n,n as b,e as R,t as B,F as N,q as F,D as y,E as L,u as o,a5 as se,a6 as K,h as A,J as le,M as ce,I as ie,N as re,a7 as de,a8 as _e,a9 as q,R as W,Q as Y,U as G}from"./index-312a31dd.js";import{i as X}from"./icon-ce016406.js";import{M as Z}from"./index-f81becf1.js";import{p as ue,d as me,a as pe}from"./project.api-bec812dd.js";import{g as ge,D as ve}from"./plugin-aa00071a.js";const he={key:0,class:"go-items-list-card"},fe={class:"list-content"},be={class:"list-content-top"},we={class:"go-flex-items-center list-footer",justify:"space-between"},ye={class:"go-flex-items-center list-footer-ri"},De=T({__name:"index",props:{cardData:Object},emits:["preview","delete","resize","edit","release"],setup(e,{emit:s}){var l;const{EllipsisHorizontalCircleSharpIcon:u,CopyIcon:m,TrashIcon:h,PencilIcon:D,DownloadIcon:d,BrowsersOutlineIcon:g,HammerIcon:t,SendIcon:_}=X.ionicons5,i=s,p=e,f=V([{label:H("global.r_edit"),key:"edit",icon:S(t)},{lable:H("global.r_more"),key:"select",icon:S(u)}]),w=j([{label:H("global.r_preview"),key:"preview",icon:S(g)},{label:(l=p.cardData)!=null&&l.release?H("global.r_unpublish"):H("global.r_publish"),key:"release",icon:S(_)},{label:H("global.r_delete"),key:"delete",icon:S(h)}]),x=v=>{switch(v){case"preview":$();break;case"delete":k();break;case"release":P();break;case"edit":E();break}},$=()=>{i("preview",p.cardData)},k=()=>{i("delete",p.cardData)},E=()=>{i("edit",p.cardData)},P=()=>{i("release",p.cardData)},I=()=>{i("resize",p.cardData)};return(v,M)=>{const U=c("n-image"),J=c("n-text"),ee=c("n-badge"),Q=c("n-button"),te=c("n-dropdown"),ae=c("n-tooltip"),ne=c("n-space"),oe=c("n-card");return e.cardData?(r(),C("div",he,[a(oe,{hoverable:"",size:"small"},{action:n(()=>[b("div",we,[a(J,{class:"go-ellipsis-1"},{default:n(()=>[R(B(e.cardData.title||e.cardData.id||"未命名"),1)]),_:1}),b("div",ye,[a(ne,null,{default:n(()=>[a(J,null,{default:n(()=>[a(ee,{class:"go-animation-twinkle",dot:"",color:e.cardData.release?"#34c749":"#fcbc40"},null,8,["color"]),R(" "+B(e.cardData.release?v.$t("project.release"):v.$t("project.unreleased")),1)]),_:1}),(r(!0),C(N,null,F(f,z=>(r(),C(N,{key:z.key},[z.key==="select"?(r(),y(te,{key:0,trigger:"hover",placement:"bottom",options:w.value,"show-arrow":!0,onSelect:x},{default:n(()=>[a(Q,{size:"small"},{icon:n(()=>[(r(),y(L(z.icon)))]),_:2},1024)]),_:2},1032,["options"])):(r(),y(ae,{key:1,placement:"bottom",trigger:"hover"},{trigger:n(()=>[a(Q,{size:"small",onClick:Le=>x(z.key)},{icon:n(()=>[(r(),y(L(z.icon)))]),_:2},1032,["onClick"])]),default:n(()=>[(r(),y(L(z.label)))]),_:2},1024))],64))),128))]),_:1})])])]),default:n(()=>[b("div",fe,[b("div",be,[a(o(Z),{class:"top-btn",hidden:["remove"],onClose:k,onResize:I})]),b("div",{class:"list-content-img",onClick:I},[a(U,{"object-fit":"contain",height:"180","preview-disabled":"",src:`${e.cardData.image}?time=${new Date().getTime()}`,alt:e.cardData.title,"fallback-src":o(se)()},null,8,["src","alt","fallback-src"])])])]),_:1})])):K("",!0)}}});const ke=A(De,[["__scopeId","data-v-0814e2b8"]]),xe={class:"list-content"},He={class:"list-content-img"},Ce=["src","alt"],$e=T({__name:"index",props:{modalShow:{required:!0,type:Boolean},cardData:{required:!0,type:Object}},emits:["close","edit"],setup(e,{emit:s}){const{HammerIcon:u}=X.ionicons5,m=j(!1),h=s,D=e;le(()=>D.modalShow,i=>{m.value=i},{immediate:!0});const d=V([{label:H("global.r_edit"),key:"edit",icon:S(u)}]),g=i=>{switch(i){case"edit":t();break}},t=()=>{h("edit",D.cardData)},_=()=>{h("close")};return(i,p)=>{const f=c("n-text"),w=c("n-space"),x=c("n-time"),$=c("n-badge"),k=c("n-button"),E=c("n-tooltip"),P=c("n-card"),I=c("n-modal");return r(),y(I,{class:"go-modal-box",show:m.value,"onUpdate:show":p[0]||(p[0]=l=>m.value=l),onAfterLeave:_},{default:n(()=>[a(P,{hoverable:"",size:"small"},{action:n(()=>[a(w,{class:"list-footer",justify:"space-between"},{default:n(()=>[a(f,{depth:"3"},{default:n(()=>[R(B(i.$t("project.last_edit"))+": ",1),a(x,{time:new Date,format:"yyyy-MM-dd hh:mm"},null,8,["time"])]),_:1}),a(w,null,{default:n(()=>[a(f,null,{default:n(()=>{var l,v;return[a($,{class:"go-animation-twinkle",dot:"",color:(l=e.cardData)!=null&&l.release?"#34c749":"#fcbc40"},null,8,["color"]),R(" "+B((v=e.cardData)!=null&&v.release?i.$t("project.release"):i.$t("project.unreleased")),1)]}),_:1}),(r(!0),C(N,null,F(d,l=>(r(),y(E,{key:l.key,placement:"bottom",trigger:"hover"},{trigger:n(()=>[a(k,{size:"small",onClick:v=>g(l.key)},{icon:n(()=>[(r(),y(L(l.icon)))]),_:2},1032,["onClick"])]),default:n(()=>[(r(),y(L(l.label)))]),_:2},1024))),128))]),_:1})]),_:1})]),default:n(()=>{var l,v;return[b("div",xe,[a(w,{class:"list-content-top go-px-0",justify:"center"},{default:n(()=>[a(w,null,{default:n(()=>[a(f,null,{default:n(()=>{var M,U;return[R(B(((M=e.cardData)==null?void 0:M.title)||((U=e.cardData)==null?void 0:U.id)||"未命名"),1)]}),_:1})]),_:1})]),_:1}),a(w,{class:"list-content-top"},{default:n(()=>[a(o(Z),{narrow:!0,hidden:["close"],onRemove:_})]),_:1}),b("div",He,[b("img",{src:(l=e.cardData)==null?void 0:l.image,alt:(v=e.cardData)==null?void 0:v.title},null,8,Ce)])])]}),_:1})]),_:1},8,["show"])}}});const Ie=A($e,[["__scopeId","data-v-0b024926"]]),ze=()=>{const e=j(!1),s=j(null);return{modalData:s,modalShow:e,closeModal:()=>{e.value=!1,s.value=null},resizeHandle:d=>{d&&(e.value=!0,s.value=d)},editHandle:d=>{if(!d)return;const g=ce(ie.CHART_HOME_NAME,"href");re(g,[d.id],void 0,!0)},previewHandle:d=>{de(_e(d.id))}}},Se=()=>{const e=j(!0),s=V({page:1,limit:12,count:10}),u=j([]),m=()=>O(void 0,null,function*(){e.value=!0;const t=yield ue({page:s.page,limit:s.limit});if(t&&t.data){const{count:_}=t;s.count=_,u.value=t.data.map(i=>{const{id:p,projectName:f,state:w,createTime:x,indexImage:$,createUserId:k}=i;return{id:p,title:f,createId:k,time:x,image:$,release:w!==-1}}),setTimeout(()=>{e.value=!1},500);return}q()}),h=t=>{s.page=t,m()},D=t=>{s.limit=t,m()},d=t=>{ge({type:ve.DELETE,promise:!0,onPositiveCallback:()=>new Promise(_=>{_(me({ids:t.id}))}),promiseResCallback:_=>{if(_.code===W.SUCCESS){window.$message.success(window.$t("global.r_delete_success")),m();return}q()}})},g=(t,_)=>O(void 0,null,function*(){const{id:i,release:p}=t,f=yield pe({id:i,state:p?-1:1});if(f&&f.code===W.SUCCESS){if(u.value=[],m(),p){window.$message.success(window.$t("global.r_unpublish_success"));return}window.$message.success(window.$t("global.r_publish_success"));return}q()});return m(),{loading:e,paginat:s,list:u,fetchList:m,releaseHandle:g,changeSize:D,changePage:h,deleteHandle:d}},je={class:"go-items-list"},Ee={class:"list-pagination"},Pe=T({__name:"index",setup(e){const{modalData:s,modalShow:u,closeModal:m,previewHandle:h,resizeHandle:D,editHandle:d}=ze(),{loading:g,paginat:t,list:_,changeSize:i,changePage:p,releaseHandle:f,deleteHandle:w}=Se();return(x,$)=>{const k=c("go-loading"),E=c("n-grid-item"),P=c("n-grid"),I=c("n-pagination");return r(),C(N,null,[b("div",je,[Y(b("div",null,[a(k)],512),[[G,o(g)]]),Y(b("div",null,[a(P,{"x-gap":20,"y-gap":20,cols:"2 s:2 m:3 l:4 xl:4 xxl:4",responsive:"screen"},{default:n(()=>[(r(!0),C(N,null,F(o(_),(l,v)=>(r(),y(E,{key:l.id},{default:n(()=>[a(o(ke),{cardData:l,onPreview:o(h),onResize:o(D),onDelete:M=>o(w)(l),onRelease:M=>o(f)(l,v),onEdit:o(d)},null,8,["cardData","onPreview","onResize","onDelete","onRelease","onEdit"])]),_:2},1024))),128))]),_:1})],512),[[G,!o(g)]]),b("div",Ee,[a(I,{page:o(t).page,"page-size":o(t).limit,"item-count":o(t).count,"page-sizes":[12,24,36,48],"onUpdate:page":o(p),"onUpdate:pageSize":o(i),"show-size-picker":""},null,8,["page","page-size","item-count","onUpdate:page","onUpdate:pageSize"])])]),o(s)?(r(),y(o(Ie),{key:0,modalShow:o(u),cardData:o(s),onClose:o(m),onEdit:o(d)},null,8,["modalShow","cardData","onClose","onEdit"])):K("",!0)],64)}}});const Me=A(Pe,[["__scopeId","data-v-c7933e27"]]),Re={class:"go-project-items"},Be=T({__name:"index",setup(e){return(s,u)=>(r(),C("div",Re,[a(o(Me))]))}});const Ve=A(Be,[["__scopeId","data-v-ba0e1d55"]]);export{Ve as default};
