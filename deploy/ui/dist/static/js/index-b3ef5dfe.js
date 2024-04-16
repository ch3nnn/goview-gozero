var Y=Object.defineProperty;var G=Object.getOwnPropertySymbols;var K=Object.prototype.hasOwnProperty,Q=Object.prototype.propertyIsEnumerable;var j=(a,s,t)=>s in a?Y(a,s,{enumerable:!0,configurable:!0,writable:!0,value:t}):a[s]=t,E=(a,s)=>{for(var t in s||(s={}))K.call(s,t)&&j(a,t,s[t]);if(G)for(var t of G(s))Q.call(s,t)&&j(a,t,s[t]);return a};var N=(a,s,t)=>new Promise((I,w)=>{var m=i=>{try{f(t.next(i))}catch(v){w(v)}},R=i=>{try{f(t.throw(i))}catch(v){w(v)}},f=i=>i.done?I(i.value):Promise.resolve(i.value).then(m,R);f((t=t.apply(a,s)).next())});import{d as X,j as k,a8 as D,B as Z,b9 as ee,a1 as te,r as _,o as S,c as V,b as o,w as e,F as $,q as oe,D as z,E as se,n as ne,t as A,e as h,u as M,aM as ae,R as re,a9 as ce,Y as x,M as le,ba as ie,aj as ue,ak as C,aJ as L,N as pe,h as de}from"./index-312a31dd.js";import{u as _e}from"./index-d277f175.js";import{u as me,P as fe}from"./chartEditStore-32c50897.js";import{s as ve}from"./useSyncUpdate.hook-4b2b135f.js";import{a as ye}from"./project.api-bec812dd.js";import{i as he}from"./icon-ce016406.js";import{l as we}from"./index-5e1011db.js";import"./plugin-aa00071a.js";import"./useSync.hook-4d532843.js";import"./chartLayoutStore-a3dfb071.js";import"./SettingItem-dba63559.js";/* empty css                                                                      */import"./SettingItemBox-19ac3d74.js";import"./CollapseItem.vue_vue_type_script_setup_true_lang-da631111.js";import"./index.esm.min-f20ed2aa.js";import"./fileTypeEnum-21359a08.js";const ge=X({__name:"index",setup(a){const{BrowsersOutlineIcon:s,SendIcon:t,AnalyticsIcon:I,CloseIcon:w}=he.ionicons5,m=me(),R=k(D()),{copy:f,isSupported:i}=_e({source:R}),v=Z(),y=k(!1),p=k(!1);ee(()=>{p.value=m.getProjectInfo.release||!1});const b=()=>{y.value=!1},F=()=>{const n=le(ie.CHART_PREVIEW_NAME,"href");if(!n)return;const{id:r}=v.params,c=typeof r=="string"?r:r[0],u=m.getStorageInfo(),l=ue(C.GO_CHART_STORAGE_LIST)||[];if(l!=null&&l.length){const g=l.findIndex(T=>T.id===c);g!==-1?(l.splice(g,1,E({id:c},u)),L(C.GO_CHART_STORAGE_LIST,l)):(l.push(E({id:c},u)),L(C.GO_CHART_STORAGE_LIST,l))}else L(C.GO_CHART_STORAGE_LIST,[E({id:c},u)]);pe(n,[c],void 0,!0)},P=()=>{y.value=!y.value},H=(n,r)=>{i?(f(),window.$message.success(n||"复制成功！")):window.$message.error(r||"复制失败！")},U=()=>N(this,null,function*(){const n=yield ye({id:ae(),state:p.value?-1:1});n&&n.code===re.SUCCESS?(P(),p.value?window.$message.success("已取消发布"):H("发布成功！已复制地址到剪贴板~","发布成功！"),m.setProjectInfo(fe.RELEASE,!p.value)):ce()}),B=[{select:!0,title:()=>"同步内容",type:()=>"primary",icon:x(I),event:ve},{key:"preview",title:()=>"预览",type:()=>"default",icon:x(s),event:F},{key:"release",title:()=>p.value?"已发布":"发布",icon:x(t),type:()=>p.value?"primary":"default",event:P}],q=te(()=>{if(m.getEditCanvas.isCodeEdit)return B;const n=we.cloneDeep(B);return n.shift(),n});return(n,r)=>{const c=_("n-button"),u=_("n-space"),l=_("n-h3"),g=_("n-icon"),T=_("n-alert"),O=_("n-list-item"),J=_("n-list"),W=_("n-modal");return S(),V($,null,[o(u,{class:"go-mt-0"},{default:e(()=>[(S(!0),V($,null,oe(q.value,d=>(S(),z(c,{key:d.key,type:d.type(),ghost:"",onClick:d.event},{icon:e(()=>[(S(),z(se(d.icon)))]),default:e(()=>[ne("span",null,A(d.title()),1)]),_:2},1032,["type","onClick"]))),128))]),_:1}),o(W,{show:y.value,"onUpdate:show":r[1]||(r[1]=d=>y.value=d),onAfterLeave:b},{default:e(()=>[o(J,{bordered:"",class:"go-system-setting"},{header:e(()=>[o(u,{justify:"space-between"},{default:e(()=>[o(l,{class:"go-mb-0"},{default:e(()=>[h("发布管理")]),_:1}),o(g,{size:"20",class:"go-cursor-pointer",onClick:b},{default:e(()=>[o(M(w))]),_:1})]),_:1})]),default:e(()=>[o(O,null,{default:e(()=>[o(u,{size:10},{default:e(()=>[o(T,{"show-icon":!1,title:"预览地址：",type:"success"},{default:e(()=>[h(A(M(D)()),1)]),_:1}),o(u,{vertical:""},{default:e(()=>[o(c,{tertiary:"",type:"primary",onClick:r[0]||(r[0]=d=>H())},{default:e(()=>[h(" 复制地址 ")]),_:1}),o(c,{type:p.value?"warning":"primary",onClick:U},{default:e(()=>[h(A(p.value?"取消发布":"发布大屏"),1)]),_:1},8,["type"])]),_:1})]),_:1})]),_:1}),o(O,null,{default:e(()=>[o(u,{size:10},{default:e(()=>[o(c,{onClick:P},{default:e(()=>[h("关闭弹窗")]),_:1})]),_:1})]),_:1})]),_:1})]),_:1},8,["show"])],64)}}});const Ne=de(ge,[["__scopeId","data-v-d4d41241"]]);export{Ne as default};