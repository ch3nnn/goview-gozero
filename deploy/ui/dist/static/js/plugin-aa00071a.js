var T=Object.defineProperty,E=Object.defineProperties;var I=Object.getOwnPropertyDescriptors;var f=Object.getOwnPropertySymbols;var S=Object.prototype.hasOwnProperty,$=Object.prototype.propertyIsEnumerable;var v=(e,o,i)=>o in e?T(e,o,{enumerable:!0,configurable:!0,writable:!0,value:i}):e[o]=i,N=(e,o)=>{for(var i in o||(o={}))S.call(o,i)&&v(e,i,o[i]);if(f)for(var i of f(o))$.call(o,i)&&v(e,i,o[i]);return e},k=(e,o)=>E(e,I(o));var m=(e,o,i)=>new Promise((n,l)=>{var g=s=>{try{r(i.next(s))}catch(t){l(t)}},d=s=>{try{r(i.throw(s))}catch(t){l(t)}},r=s=>s.done?n(s.value):Promise.resolve(s.value).then(g,d);r((i=i.apply(e,o)).next())});import{Y as x,b6 as y,b7 as u}from"./index-312a31dd.js";import{i as A}from"./icon-ce016406.js";var a=(e=>(e.DELETE="delete",e.WARNING="warning",e.ERROR="error",e.SUCCESS="success",e))(a||{});const{InformationCircleIcon:G}=A.ionicons5,j=()=>{window.$loading.start()},z=()=>{setTimeout(()=>{window.$loading.finish()})},L=()=>{setTimeout(()=>{window.$loading.error()})},P=e=>{const{type:o,title:i,message:n,positiveText:l,negativeText:g,closeNegativeText:d,isMaskClosable:r,onPositiveCallback:s,onNegativeCallback:t,promise:p,promiseResCallback:C,promiseRejCallback:R}=e,b={[a.DELETE]:{fn:window.$dialog.warning,message:n||"是否删除此数据?"},[a.WARNING]:{fn:window.$dialog.warning,message:n||"是否执行此操作?"},[a.ERROR]:{fn:window.$dialog.error,message:n||"是否执行此操作?"},[a.SUCCESS]:{fn:window.$dialog.success,message:n||"是否执行此操作?"}},c=b[o||a.WARNING].fn(k(N({},e),{title:i||"提示",icon:x(G,{size:u}),content:b[o||a.WARNING].message,positiveText:l||"确定",negativeText:d?void 0:g||"取消",maskClosable:r||y,onPositiveClick:()=>m(void 0,null,function*(){if(p&&s){c.loading=!0;try{const w=yield s();C&&C(w)}catch(w){R&&R(w)}c.loading=!1;return}s&&s(c)}),onNegativeClick:()=>m(void 0,null,function*(){t&&t(c)})}))};export{a as D,j as a,z as b,P as g,L as l};
