var z=(l,s,r)=>new Promise((c,_)=>{var p=d=>{try{f(r.next(d))}catch(i){_(i)}},m=d=>{try{f(r.throw(d))}catch(i){_(i)}},f=d=>d.done?c(d.value):Promise.resolve(d.value).then(p,m);f((r=r.apply(l,s)).next())});import{d as S,j as P,G as F,H as N,I as D,P as E,J as G,r as n,o as a,D as u,w as e,b as t,e as x,t as k,E as C,u as o,c as B,q,F as K,L as J,R as Q,M as X,N as Y,h as I,O as Z,n as w,Q as j,U as T,V as ee,W as te,l as oe,X as A,Y as O,Z as L,$ as ne,a0 as ae,B as se,a1 as ce,m as le,a2 as re,a3 as _e,a4 as ie,K as de}from"./index-312a31dd.js";import{i as M}from"./icon-ce016406.js";import{c as ue}from"./project.api-bec812dd.js";import{_ as pe}from"./index.vue_vue_type_script_setup_true_lang-cdfd7702.js";import{g as me}from"./plugin-aa00071a.js";import"./index-d277f175.js";const fe=S({__name:"index",props:{show:Boolean},emits:["close"],setup(l,{emit:s}){const{FishIcon:r,CloseIcon:c}=M.ionicons5,{StoreIcon:_,ObjectStorageIcon:p}=M.carbon,m=P(!1),f=s,d=l,i=F([{title:N("project.new_project"),key:D.CHART_HOME_NAME,icon:r,disabled:!1},{title:N("project.my_templete"),key:E.BASE_HOME_TEMPLATE_NAME,icon:p,disabled:!0},{title:N("project.template_market"),key:E.BASE_HOME_TEMPLATE_MARKET_NAME,icon:_,disabled:!0}]);G(()=>d.show,y=>{m.value=y});const v=()=>{f("close",!1)},g=y=>z(this,null,function*(){switch(y){case D.CHART_HOME_NAME:try{const h=yield ue({projectName:J(),remarks:null,indexImage:null});if(h&&h.code===Q.SUCCESS){window.$message.success(window.$t("project.create_success"));const{id:$}=h.data,H=X(D.CHART_HOME_NAME,"href");Y(H,[$],void 0,!0),v()}}catch(h){window.$message.error(window.$t("project.create_failure"))}break}});return(y,h)=>{const $=n("n-text"),H=n("n-icon"),W=n("n-button"),R=n("n-space"),U=n("n-card"),V=n("n-modal");return a(),u(V,{show:m.value,"onUpdate:show":h[0]||(h[0]=b=>m.value=b),class:"go-create-modal",onAfterLeave:v},{default:e(()=>[t(R,{size:"large"},{default:e(()=>[t(U,{class:"card-box",hoverable:""},{header:e(()=>[t($,{class:"card-box-tite"},{default:e(()=>[x(k(y.$t("project.create_tip")),1)]),_:1})]),"header-extra":e(()=>[t($,{onClick:v},{default:e(()=>[t(H,{size:"20"},{default:e(()=>[(a(),u(C(o(c))))]),_:1})]),_:1})]),action:e(()=>[]),default:e(()=>[t(R,{class:"card-box-content",justify:"center"},{default:e(()=>[(a(!0),B(K,null,q(i.value,b=>(a(),u(W,{size:"large",disabled:b.disabled,key:b.key,onClick:Ne=>g(b.key)},{icon:e(()=>[t(H,{size:"18"},{default:e(()=>[(a(),u(C(b.icon)))]),_:2},1024)]),default:e(()=>[(a(),u(C(b.title)))]),_:2},1032,["disabled","onClick"]))),128))]),_:1})]),_:1})]),_:1})]),_:1},8,["show"])}}});const ve=I(fe,[["__scopeId","data-v-7ac7d745"]]),ye=S({__name:"index",props:{collapsed:Boolean},setup(l){const{DuplicateIcon:s,DuplicateOutlineIcon:r}=M.ionicons5,c=Z(),_=P(!1),p=()=>{_.value=!0},m=()=>{_.value=!1};return(f,d)=>{const i=n("n-icon"),v=n("n-button"),g=n("n-tooltip");return a(),B(K,null,[w("div",{onClick:p},[l.collapsed?(a(),u(g,{key:0,placement:"right",trigger:"hover"},{trigger:e(()=>[t(v,{ghost:"",type:"primary",size:"small"},{icon:e(()=>[t(i,null,{default:e(()=>[j(t(o(r),null,null,512),[[T,o(c).getDarkTheme]]),j(t(o(s),null,null,512),[[T,!o(c).getDarkTheme]])]),_:1})]),_:1})]),default:e(()=>[w("span",null,k(f.$t("project.create_btn")),1)]),_:1})):(a(),u(v,{key:1,ghost:"",type:"primary"},{icon:e(()=>[t(i,null,{default:e(()=>[j(t(o(r),null,null,512),[[T,o(c).getDarkTheme]]),j(t(o(s),null,null,512),[[T,!o(c).getDarkTheme]])]),_:1})]),default:e(()=>[w("span",null,k(f.$t("project.create_btn")),1)]),_:1}))]),t(o(ve),{show:_.value,onClose:m},null,8,["show"])],64)}}}),he={class:"go-aside-footer"},ge=S({__name:"index",props:{collapsed:Boolean},setup(l){const{DocumentTextIcon:s,CodeSlashIcon:r}=M.ionicons5,c=()=>{ee()},_=()=>{te()};return(p,m)=>{const f=n("n-divider"),d=n("n-icon"),i=n("n-button"),v=n("n-text"),g=n("n-tooltip"),y=n("n-space");return a(),B("div",he,[t(f,{class:"go-mt-0"}),t(y,{justify:"space-around"},{default:e(()=>[l.collapsed?(a(),u(g,{key:0,placement:"right",trigger:"hover"},{trigger:e(()=>[t(i,{secondary:"",onClick:c},{icon:e(()=>[t(d,{size:"18"},{default:e(()=>[t(o(s))]),_:1})]),_:1})]),default:e(()=>[t(v,null,{default:e(()=>[x(k(p.$t("global.doc")),1)]),_:1})]),_:1})):(a(),u(i,{key:1,secondary:"",onClick:c},{icon:e(()=>[t(d,{size:"18"},{default:e(()=>[t(o(s))]),_:1})]),default:e(()=>[t(v,null,{default:e(()=>[x(k(p.$t("global.doc")),1)]),_:1})]),_:1})),l.collapsed?(a(),u(g,{key:2,placement:"right",trigger:"hover"},{trigger:e(()=>[t(i,{secondary:"",onClick:c},{icon:e(()=>[t(d,{size:"18"},{default:e(()=>[t(o(r))]),_:1})]),_:1})]),default:e(()=>[t(v,null,{default:e(()=>[x(k(p.$t("global.code_addr")),1)]),_:1})]),_:1})):(a(),u(i,{key:3,secondary:"",onClick:_},{icon:e(()=>[t(d,{size:"18"},{default:e(()=>[t(o(r))]),_:1})]),default:e(()=>[j(t(v,null,{default:e(()=>[x(k(p.$t("global.code_addr")),1)]),_:1},512),[[T,!l.collapsed]])]),_:1}))]),_:1})])}}});const Ee=I(ge,[["__scopeId","data-v-b5a02cb8"]]),{GridIcon:Ue,TvOutlineIcon:be}=M.ionicons5,{StoreIcon:ke,ObjectStorageIcon:we,DevicesIcon:Me}=M.carbon,Ae=()=>["all-project"],xe=()=>{const l=window.$t;return oe([{key:"divider-1",type:"divider"},{label:()=>A("span",null,{default:()=>l("project.project")}),key:"all-project",icon:O(Me),children:[{type:"group",label:()=>A("span",null,{default:()=>l("project.my")}),key:"my-project",children:[{label:()=>A(L,{to:{name:E.BASE_HOME_ITEMS_NAME}},{default:()=>l("project.all_project")}),key:E.BASE_HOME_ITEMS_NAME,icon:O(be)},{label:()=>A(L,{to:{name:E.BASE_HOME_TEMPLATE_NAME}},{default:()=>l("project.my_templete")}),key:E.BASE_HOME_TEMPLATE_NAME,icon:O(we)}]}]},{key:"divider-2",type:"divider"},{label:()=>A(L,{to:{name:E.BASE_HOME_TEMPLATE_MARKET_NAME}},{default:()=>l("project.template_market")}),key:E.BASE_HOME_TEMPLATE_MARKET_NAME,icon:O(ke)}])},je={class:"go-project-sider-flex"},Te={class:"sider-bottom"},Ce=S({__name:"index",setup(l){const s=P(!1),{getAsideCollapsedWidth:r}=ne(ae()),c=se(),_=ce(()=>c.name),p=xe(),m=Ae(),f=()=>{document.body.clientWidth<=950?s.value=!0:s.value=!1};return le(()=>{window.addEventListener("resize",f)}),re(()=>{window.removeEventListener("resize",f)}),(d,i)=>{const v=n("n-space"),g=n("n-menu"),y=n("n-layout-sider");return a(),u(y,{class:"go-project-sider",bordered:"","collapse-mode":"width","show-trigger":"bar",collapsed:s.value,"native-scrollbar":!1,"collapsed-width":o(r),width:o(_e),onCollapse:i[0]||(i[0]=h=>s.value=!0),onExpand:i[1]||(i[1]=h=>s.value=!1)},{default:e(()=>[w("div",je,[w("aside",null,[t(v,{vertical:"",class:"go-project-sider-top"},{default:e(()=>[t(o(ye),{collapsed:s.value},null,8,["collapsed"])]),_:1}),t(g,{value:_.value,options:o(p),"collapsed-width":o(r),"collapsed-icon-size":22,"default-expanded-keys":o(m)},null,8,["value","options","collapsed-width","default-expanded-keys"])]),w("div",Te,[t(o(Ee),{collapsed:s.value},null,8,["collapsed"])])])]),_:1},8,["collapsed","collapsed-width","width"])}}});const Se=I(Ce,[["__scopeId","data-v-58606bc9"]]),Ie={};function $e(l,s){const r=n("router-view");return a(),u(r,null,{default:e(({Component:c,route:_})=>[t(ie,{name:"fade",mode:"out-in",appear:""},{default:e(()=>[_.meta.noKeepAlive?(a(),u(C(c),{key:_.fullPath})):(a(),u(de,{key:1},[(a(),u(C(c),{key:_.fullPath}))],1024))]),_:2},1024)]),_:1})}const He=I(Ie,[["render",$e]]),Oe={class:"go-project"},Be=S({__name:"index",setup(l){return me({message:"不要在官方后端上发布任何私密数据，任何人都看得到并进行删除！！！！",isMaskClosable:!0,closeNegativeText:!0,transformOrigin:"center",onPositiveCallback:()=>{}}),(s,r)=>{const c=n("n-space"),_=n("router-view"),p=n("n-layout-content"),m=n("n-layout");return a(),B("div",Oe,[t(m,{"has-sider":"",position:"absolute"},{default:e(()=>[t(c,{vertical:""},{default:e(()=>[t(o(Se))]),_:1}),t(m,null,{default:e(()=>[t(o(pe)),t(m,{id:"go-project-content-top",class:"content-top",position:"absolute","native-scrollbar":!1},{default:e(()=>[t(p,null,{default:e(()=>[t(o(He),null,{default:e(()=>[t(_)]),_:1})]),_:1})]),_:1})]),_:1})]),_:1})])}}});const Ve=I(Be,[["__scopeId","data-v-59301cbb"]]);export{Ve as default};