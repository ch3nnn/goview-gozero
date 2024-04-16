var d=Object.defineProperty;var i=Object.getOwnPropertySymbols;var O=Object.prototype.hasOwnProperty,f=Object.prototype.propertyIsEnumerable;var E=(e,n,t)=>n in e?d(e,n,{enumerable:!0,configurable:!0,writable:!0,value:t}):e[n]=t,c=(e,n)=>{for(var t in n||(n={}))O.call(n,t)&&E(e,t,n[t]);if(i)for(var t of i(n))f.call(n,t)&&E(e,t,n[t]);return e};import{aW as r,cy as p}from"./index-312a31dd.js";const u={},a={echarts:p},b=e=>{if(!e.events)return{};const n={};for(const o in e.events.baseEvent){const s=e.events.baseEvent[o];s&&(n[o]=N(s))}const t=e.events.advancedEvents||{},m={[r.VNODE_BEFORE_MOUNT](o){u[e.id]=o.component;const s=(t[r.VNODE_BEFORE_MOUNT]||"").trim();v(s,o)},[r.VNODE_MOUNTED](o){const s=(t[r.VNODE_MOUNTED]||"").trim();v(s,o)}};return c(c({},n),m)};function N(e){try{return new Function(`
      return (
        async function(components,mouseEvent){
          ${e}
        }
      )`)().bind(void 0,u)}catch(n){console.error(n)}}function v(e,n){try{Function(`
      "use strict";
      return (
        async function(e, components, node_modules){
          const {${Object.keys(a).join()}} = node_modules;
          ${e}
        }
      )`)().bind(n==null?void 0:n.component)(n,u,a)}catch(t){console.error(t)}}export{a as n,b as u};
