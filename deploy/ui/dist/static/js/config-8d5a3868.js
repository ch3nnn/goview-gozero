var s=Object.defineProperty,m=Object.defineProperties;var d=Object.getOwnPropertyDescriptors;var r=Object.getOwnPropertySymbols;var g=Object.prototype.hasOwnProperty,C=Object.prototype.propertyIsEnumerable;var i=(p,e,o)=>e in p?s(p,e,{enumerable:!0,configurable:!0,writable:!0,value:o}):p[e]=o,a=(p,e)=>{for(var o in e||(e={}))g.call(e,o)&&i(p,o,e[o]);if(r)for(var o of r(e))C.call(e,o)&&i(p,o,e[o]);return p},f=(p,e)=>m(p,d(e));var t=(p,e,o)=>(i(p,typeof e!="symbol"?e+"":e,o),o);import{aU as c,ac as l}from"./index-312a31dd.js";import{e as h}from"./chartEditStore-32c50897.js";import{g as n}from"./index-5e1011db.js";import"./plugin-aa00071a.js";import"./icon-ce016406.js";import"./SettingItem-dba63559.js";/* empty css                                                                      */import"./SettingItemBox-19ac3d74.js";import"./CollapseItem.vue_vue_type_script_setup_true_lang-da631111.js";import"./index.esm.min-f20ed2aa.js";import"./fileTypeEnum-21359a08.js";const D={dataset:10*60,useEndDate:!1,endDate:new Date().getTime(),style:"时分秒",showDay:!1,flipperBgColor:"#16293E",flipperTextColor:"#4A9EF8FF",flipperWidth:30,flipperHeight:50,flipperRadius:5,flipperGap:10,flipperType:"down",flipperSpeed:450};class B extends h{constructor(){super(...arguments);t(this,"key",n.key);t(this,"attr",f(a({},c),{w:500,h:100,zIndex:-1}));t(this,"chartConfig",l(n));t(this,"option",l(D))}}export{B as default,D as option};
