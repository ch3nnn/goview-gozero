var H=Object.defineProperty,V=Object.defineProperties;var W=Object.getOwnPropertyDescriptors;var O=Object.getOwnPropertySymbols;var j=Object.prototype.hasOwnProperty,B=Object.prototype.propertyIsEnumerable;var A=(e,t,r)=>t in e?H(e,t,{enumerable:!0,configurable:!0,writable:!0,value:r}):e[t]=r,I=(e,t)=>{for(var r in t||(t={}))j.call(t,r)&&A(e,r,t[r]);if(O)for(var r of O(t))B.call(t,r)&&A(e,r,t[r]);return e},N=(e,t)=>V(e,W(t));var w=(e,t,r)=>(A(e,typeof t!="symbol"?t+"":t,r),r);var m=(e,t,r)=>new Promise((i,a)=>{var g=s=>{try{u(r.next(s))}catch(d){a(d)}},c=s=>{try{u(r.throw(s))}catch(d){a(d)}},u=s=>s.done?i(s.value):Promise.resolve(s.value).then(g,c);u((r=r.apply(e,t)).next())});import{t as U,m as $}from"./EditorWorker-4b89d885.js";import"./index-312a31dd.js";import"./editorWorker-43a98755.js";/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.33.0(4b1abad427e58dbedc1215d99a0902ffc885fcd4)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/var M=Object.defineProperty,z=Object.getOwnPropertyDescriptor,G=Object.getOwnPropertyNames,J=Object.prototype.hasOwnProperty,Q=(e,t,r)=>t in e?M(e,t,{enumerable:!0,configurable:!0,writable:!0,value:r}):e[t]=r,q=(e,t,r,i)=>{if(t&&typeof t=="object"||typeof t=="function")for(let a of G(t))!J.call(e,a)&&(r||a!=="default")&&M(e,a,{get:()=>t[a],enumerable:!(i=z(t,a))||i.enumerable});return e},h=(e,t,r)=>(Q(e,typeof t!="symbol"?t+"":t,r),r),n={};q(n,$);var X=class{constructor(e,t){w(this,"_modeId");w(this,"_defaults");w(this,"_configChangeListener");w(this,"_updateExtraLibsToken");w(this,"_extraLibsChangeListener");w(this,"_worker");w(this,"_client");this._modeId=e,this._defaults=t,this._worker=null,this._client=null,this._configChangeListener=this._defaults.onDidChange(()=>this._stopWorker()),this._updateExtraLibsToken=0,this._extraLibsChangeListener=this._defaults.onDidExtraLibsChange(()=>this._updateExtraLibs())}_stopWorker(){this._worker&&(this._worker.dispose(),this._worker=null),this._client=null}dispose(){this._configChangeListener.dispose(),this._extraLibsChangeListener.dispose(),this._stopWorker()}_updateExtraLibs(){return m(this,null,function*(){if(!this._worker)return;const e=++this._updateExtraLibsToken,t=yield this._worker.getProxy();this._updateExtraLibsToken===e&&t.updateExtraLibs(this._defaults.getExtraLibs())})}_getClient(){if(!this._client){this._worker=n.editor.createWebWorker({moduleId:"vs/language/typescript/tsWorker",label:this._modeId,keepIdleModels:!0,createData:{compilerOptions:this._defaults.getCompilerOptions(),extraLibs:this._defaults.getExtraLibs(),customWorkerPath:this._defaults.workerOptions.customWorkerPath,inlayHintsOptions:this._defaults.inlayHintsOptions}});let e=this._worker.getProxy();this._defaults.getEagerModelSync()&&(e=e.then(t=>this._worker?this._worker.withSyncedResources(n.editor.getModels().filter(r=>r.getLanguageId()===this._modeId).map(r=>r.uri)):t)),this._client=e}return this._client}getLanguageServiceWorker(...e){let t;return this._getClient().then(r=>{t=r}).then(r=>{if(this._worker)return this._worker.withSyncedResources(e)}).then(r=>t)}},o={};o["lib.d.ts"]=!0;o["lib.dom.d.ts"]=!0;o["lib.dom.iterable.d.ts"]=!0;o["lib.es2015.collection.d.ts"]=!0;o["lib.es2015.core.d.ts"]=!0;o["lib.es2015.d.ts"]=!0;o["lib.es2015.generator.d.ts"]=!0;o["lib.es2015.iterable.d.ts"]=!0;o["lib.es2015.promise.d.ts"]=!0;o["lib.es2015.proxy.d.ts"]=!0;o["lib.es2015.reflect.d.ts"]=!0;o["lib.es2015.symbol.d.ts"]=!0;o["lib.es2015.symbol.wellknown.d.ts"]=!0;o["lib.es2016.array.include.d.ts"]=!0;o["lib.es2016.d.ts"]=!0;o["lib.es2016.full.d.ts"]=!0;o["lib.es2017.d.ts"]=!0;o["lib.es2017.full.d.ts"]=!0;o["lib.es2017.intl.d.ts"]=!0;o["lib.es2017.object.d.ts"]=!0;o["lib.es2017.sharedmemory.d.ts"]=!0;o["lib.es2017.string.d.ts"]=!0;o["lib.es2017.typedarrays.d.ts"]=!0;o["lib.es2018.asyncgenerator.d.ts"]=!0;o["lib.es2018.asynciterable.d.ts"]=!0;o["lib.es2018.d.ts"]=!0;o["lib.es2018.full.d.ts"]=!0;o["lib.es2018.intl.d.ts"]=!0;o["lib.es2018.promise.d.ts"]=!0;o["lib.es2018.regexp.d.ts"]=!0;o["lib.es2019.array.d.ts"]=!0;o["lib.es2019.d.ts"]=!0;o["lib.es2019.full.d.ts"]=!0;o["lib.es2019.object.d.ts"]=!0;o["lib.es2019.string.d.ts"]=!0;o["lib.es2019.symbol.d.ts"]=!0;o["lib.es2020.bigint.d.ts"]=!0;o["lib.es2020.d.ts"]=!0;o["lib.es2020.full.d.ts"]=!0;o["lib.es2020.intl.d.ts"]=!0;o["lib.es2020.promise.d.ts"]=!0;o["lib.es2020.sharedmemory.d.ts"]=!0;o["lib.es2020.string.d.ts"]=!0;o["lib.es2020.symbol.wellknown.d.ts"]=!0;o["lib.es2021.d.ts"]=!0;o["lib.es2021.full.d.ts"]=!0;o["lib.es2021.intl.d.ts"]=!0;o["lib.es2021.promise.d.ts"]=!0;o["lib.es2021.string.d.ts"]=!0;o["lib.es2021.weakref.d.ts"]=!0;o["lib.es5.d.ts"]=!0;o["lib.es6.d.ts"]=!0;o["lib.esnext.d.ts"]=!0;o["lib.esnext.full.d.ts"]=!0;o["lib.esnext.intl.d.ts"]=!0;o["lib.esnext.promise.d.ts"]=!0;o["lib.esnext.string.d.ts"]=!0;o["lib.esnext.weakref.d.ts"]=!0;o["lib.scripthost.d.ts"]=!0;o["lib.webworker.d.ts"]=!0;o["lib.webworker.importscripts.d.ts"]=!0;o["lib.webworker.iterable.d.ts"]=!0;function T(e,t,r=0){if(typeof e=="string")return e;if(e===void 0)return"";let i="";if(r){i+=t;for(let a=0;a<r;a++)i+="  "}if(i+=e.messageText,r++,e.next)for(const a of e.next)i+=T(a,t,r);return i}function S(e){return e?e.map(t=>t.text).join(""):""}var y=class{constructor(e){this._worker=e}_textSpanToRange(e,t){let r=e.getPositionAt(t.start),i=e.getPositionAt(t.start+t.length),{lineNumber:a,column:g}=r,{lineNumber:c,column:u}=i;return{startLineNumber:a,startColumn:g,endLineNumber:c,endColumn:u}}},Y=class{constructor(e){w(this,"_libFiles");w(this,"_hasFetchedLibFiles");w(this,"_fetchLibFilesPromise");this._worker=e,this._libFiles={},this._hasFetchedLibFiles=!1,this._fetchLibFilesPromise=null}isLibFile(e){return e&&e.path.indexOf("/lib.")===0?!!o[e.path.slice(1)]:!1}getOrCreateModel(e){const t=n.Uri.parse(e),r=n.editor.getModel(t);if(r)return r;if(this.isLibFile(t)&&this._hasFetchedLibFiles)return n.editor.createModel(this._libFiles[t.path.slice(1)],"typescript",t);const i=U.getExtraLibs()[e];return i?n.editor.createModel(i.content,"typescript",t):null}_containsLibFile(e){for(let t of e)if(this.isLibFile(t))return!0;return!1}fetchLibFilesIfNecessary(e){return m(this,null,function*(){this._containsLibFile(e)&&(yield this._fetchLibFiles())})}_fetchLibFiles(){return this._fetchLibFilesPromise||(this._fetchLibFilesPromise=this._worker().then(e=>e.getLibFiles()).then(e=>{this._hasFetchedLibFiles=!0,this._libFiles=e})),this._fetchLibFilesPromise}},Z=class extends y{constructor(t,r,i,a){super(a);w(this,"_disposables",[]);w(this,"_listener",Object.create(null));this._libFiles=t,this._defaults=r,this._selector=i;const g=s=>{if(s.getLanguageId()!==i)return;const d=()=>{const{onlyVisible:k}=this._defaults.getDiagnosticsOptions();k?s.isAttachedToEditor()&&this._doValidate(s):this._doValidate(s)};let p;const f=s.onDidChangeContent(()=>{clearTimeout(p),p=window.setTimeout(d,500)}),b=s.onDidChangeAttached(()=>{const{onlyVisible:k}=this._defaults.getDiagnosticsOptions();k&&(s.isAttachedToEditor()?d():n.editor.setModelMarkers(s,this._selector,[]))});this._listener[s.uri.toString()]={dispose(){f.dispose(),b.dispose(),clearTimeout(p)}},d()},c=s=>{n.editor.setModelMarkers(s,this._selector,[]);const d=s.uri.toString();this._listener[d]&&(this._listener[d].dispose(),delete this._listener[d])};this._disposables.push(n.editor.onDidCreateModel(s=>g(s))),this._disposables.push(n.editor.onWillDisposeModel(c)),this._disposables.push(n.editor.onDidChangeModelLanguage(s=>{c(s.model),g(s.model)})),this._disposables.push({dispose(){for(const s of n.editor.getModels())c(s)}});const u=()=>{for(const s of n.editor.getModels())c(s),g(s)};this._disposables.push(this._defaults.onDidChange(u)),this._disposables.push(this._defaults.onDidExtraLibsChange(u)),n.editor.getModels().forEach(s=>g(s))}dispose(){this._disposables.forEach(t=>t&&t.dispose()),this._disposables=[]}_doValidate(t){return m(this,null,function*(){const r=yield this._worker(t.uri);if(t.isDisposed())return;const i=[],{noSyntaxValidation:a,noSemanticValidation:g,noSuggestionDiagnostics:c}=this._defaults.getDiagnosticsOptions();a||i.push(r.getSyntacticDiagnostics(t.uri.toString())),g||i.push(r.getSemanticDiagnostics(t.uri.toString())),c||i.push(r.getSuggestionDiagnostics(t.uri.toString()));const u=yield Promise.all(i);if(!u||t.isDisposed())return;const s=u.reduce((p,f)=>f.concat(p),[]).filter(p=>(this._defaults.getDiagnosticsOptions().diagnosticCodesToIgnore||[]).indexOf(p.code)===-1),d=s.map(p=>p.relatedInformation||[]).reduce((p,f)=>f.concat(p),[]).map(p=>p.file?n.Uri.parse(p.file.fileName):null);yield this._libFiles.fetchLibFilesIfNecessary(d),!t.isDisposed()&&n.editor.setModelMarkers(t,this._selector,s.map(p=>this._convertDiagnostics(t,p)))})}_convertDiagnostics(t,r){const i=r.start||0,a=r.length||1,{lineNumber:g,column:c}=t.getPositionAt(i),{lineNumber:u,column:s}=t.getPositionAt(i+a),d=[];return r.reportsUnnecessary&&d.push(n.MarkerTag.Unnecessary),r.reportsDeprecated&&d.push(n.MarkerTag.Deprecated),{severity:this._tsDiagnosticCategoryToMarkerSeverity(r.category),startLineNumber:g,startColumn:c,endLineNumber:u,endColumn:s,message:T(r.messageText,`
`),code:r.code.toString(),tags:d,relatedInformation:this._convertRelatedInformation(t,r.relatedInformation)}}_convertRelatedInformation(t,r){if(!r)return[];const i=[];return r.forEach(a=>{let g=t;if(a.file&&(g=this._libFiles.getOrCreateModel(a.file.fileName)),!g)return;const c=a.start||0,u=a.length||1,{lineNumber:s,column:d}=g.getPositionAt(c),{lineNumber:p,column:f}=g.getPositionAt(c+u);i.push({resource:g.uri,startLineNumber:s,startColumn:d,endLineNumber:p,endColumn:f,message:T(a.messageText,`
`)})}),i}_tsDiagnosticCategoryToMarkerSeverity(t){switch(t){case 1:return n.MarkerSeverity.Error;case 3:return n.MarkerSeverity.Info;case 0:return n.MarkerSeverity.Warning;case 2:return n.MarkerSeverity.Hint}return n.MarkerSeverity.Info}},F=class extends y{get triggerCharacters(){return["."]}provideCompletionItems(e,t,r,i){return m(this,null,function*(){const a=e.getWordUntilPosition(t),g=new n.Range(t.lineNumber,a.startColumn,t.lineNumber,a.endColumn),c=e.uri,u=e.getOffsetAt(t),s=yield this._worker(c);if(e.isDisposed())return;const d=yield s.getCompletionsAtPosition(c.toString(),u);return!d||e.isDisposed()?void 0:{suggestions:d.entries.map(f=>{var C;let b=g;if(f.replacementSpan){const x=e.getPositionAt(f.replacementSpan.start),D=e.getPositionAt(f.replacementSpan.start+f.replacementSpan.length);b=new n.Range(x.lineNumber,x.column,D.lineNumber,D.column)}const k=[];return((C=f.kindModifiers)==null?void 0:C.indexOf("deprecated"))!==-1&&k.push(n.languages.CompletionItemTag.Deprecated),{uri:c,position:t,offset:u,range:b,label:f.name,insertText:f.name,sortText:f.sortText,kind:F.convertKind(f.kind),tags:k}})}})}resolveCompletionItem(e,t){return m(this,null,function*(){const r=e,i=r.uri,a=r.position,g=r.offset,u=yield(yield this._worker(i)).getCompletionEntryDetails(i.toString(),g,r.label);return u?{uri:i,position:a,label:u.name,kind:F.convertKind(u.kind),detail:S(u.displayParts),documentation:{value:F.createDocumentationString(u)}}:r})}static convertKind(e){switch(e){case l.primitiveType:case l.keyword:return n.languages.CompletionItemKind.Keyword;case l.variable:case l.localVariable:return n.languages.CompletionItemKind.Variable;case l.memberVariable:case l.memberGetAccessor:case l.memberSetAccessor:return n.languages.CompletionItemKind.Field;case l.function:case l.memberFunction:case l.constructSignature:case l.callSignature:case l.indexSignature:return n.languages.CompletionItemKind.Function;case l.enum:return n.languages.CompletionItemKind.Enum;case l.module:return n.languages.CompletionItemKind.Module;case l.class:return n.languages.CompletionItemKind.Class;case l.interface:return n.languages.CompletionItemKind.Interface;case l.warning:return n.languages.CompletionItemKind.File}return n.languages.CompletionItemKind.Property}static createDocumentationString(e){let t=S(e.documentation);if(e.tags)for(const r of e.tags)t+=`

${K(r)}`;return t}};function K(e){let t=`*@${e.name}*`;if(e.name==="param"&&e.text){const[r,...i]=e.text;t+=`\`${r.text}\``,i.length>0&&(t+=` — ${i.map(a=>a.text).join(" ")}`)}else Array.isArray(e.text)?t+=` — ${e.text.map(r=>r.text).join(" ")}`:e.text&&(t+=` — ${e.text}`);return t}var R=class extends y{constructor(){super(...arguments);w(this,"signatureHelpTriggerCharacters",["(",","])}static _toSignatureHelpTriggerReason(t){switch(t.triggerKind){case n.languages.SignatureHelpTriggerKind.TriggerCharacter:return t.triggerCharacter?t.isRetrigger?{kind:"retrigger",triggerCharacter:t.triggerCharacter}:{kind:"characterTyped",triggerCharacter:t.triggerCharacter}:{kind:"invoked"};case n.languages.SignatureHelpTriggerKind.ContentChange:return t.isRetrigger?{kind:"retrigger"}:{kind:"invoked"};case n.languages.SignatureHelpTriggerKind.Invoke:default:return{kind:"invoked"}}}provideSignatureHelp(t,r,i,a){return m(this,null,function*(){const g=t.uri,c=t.getOffsetAt(r),u=yield this._worker(g);if(t.isDisposed())return;const s=yield u.getSignatureHelpItems(g.toString(),c,{triggerReason:R._toSignatureHelpTriggerReason(a)});if(!s||t.isDisposed())return;const d={activeSignature:s.selectedItemIndex,activeParameter:s.argumentIndex,signatures:[]};return s.items.forEach(p=>{const f={label:"",parameters:[]};f.documentation={value:S(p.documentation)},f.label+=S(p.prefixDisplayParts),p.parameters.forEach((b,k,C)=>{const x=S(b.displayParts),D={label:x,documentation:{value:S(b.documentation)}};f.label+=x,f.parameters.push(D),k<C.length-1&&(f.label+=S(p.separatorDisplayParts))}),f.label+=S(p.suffixDisplayParts),d.signatures.push(f)}),{value:d,dispose(){}}})}},ee=class extends y{provideHover(e,t,r){return m(this,null,function*(){const i=e.uri,a=e.getOffsetAt(t),g=yield this._worker(i);if(e.isDisposed())return;const c=yield g.getQuickInfoAtPosition(i.toString(),a);if(!c||e.isDisposed())return;const u=S(c.documentation),s=c.tags?c.tags.map(p=>K(p)).join(`  

`):"",d=S(c.displayParts);return{range:this._textSpanToRange(e,c.textSpan),contents:[{value:"```typescript\n"+d+"\n```\n"},{value:u+(s?`

`+s:"")}]}})}},te=class extends y{provideDocumentHighlights(e,t,r){return m(this,null,function*(){const i=e.uri,a=e.getOffsetAt(t),g=yield this._worker(i);if(e.isDisposed())return;const c=yield g.getOccurrencesAtPosition(i.toString(),a);if(!(!c||e.isDisposed()))return c.map(u=>({range:this._textSpanToRange(e,u.textSpan),kind:u.isWriteAccess?n.languages.DocumentHighlightKind.Write:n.languages.DocumentHighlightKind.Text}))})}},re=class extends y{constructor(e,t){super(t),this._libFiles=e}provideDefinition(e,t,r){return m(this,null,function*(){const i=e.uri,a=e.getOffsetAt(t),g=yield this._worker(i);if(e.isDisposed())return;const c=yield g.getDefinitionAtPosition(i.toString(),a);if(!c||e.isDisposed()||(yield this._libFiles.fetchLibFilesIfNecessary(c.map(s=>n.Uri.parse(s.fileName))),e.isDisposed()))return;const u=[];for(let s of c){const d=this._libFiles.getOrCreateModel(s.fileName);d&&u.push({uri:d.uri,range:this._textSpanToRange(d,s.textSpan)})}return u})}},se=class extends y{constructor(e,t){super(t),this._libFiles=e}provideReferences(e,t,r,i){return m(this,null,function*(){const a=e.uri,g=e.getOffsetAt(t),c=yield this._worker(a);if(e.isDisposed())return;const u=yield c.getReferencesAtPosition(a.toString(),g);if(!u||e.isDisposed()||(yield this._libFiles.fetchLibFilesIfNecessary(u.map(d=>n.Uri.parse(d.fileName))),e.isDisposed()))return;const s=[];for(let d of u){const p=this._libFiles.getOrCreateModel(d.fileName);p&&s.push({uri:p.uri,range:this._textSpanToRange(p,d.textSpan)})}return s})}},ie=class extends y{provideDocumentSymbols(e,t){return m(this,null,function*(){const r=e.uri,i=yield this._worker(r);if(e.isDisposed())return;const a=yield i.getNavigationBarItems(r.toString());if(!a||e.isDisposed())return;const g=(u,s,d)=>{let p={name:s.text,detail:"",kind:_[s.kind]||n.languages.SymbolKind.Variable,range:this._textSpanToRange(e,s.spans[0]),selectionRange:this._textSpanToRange(e,s.spans[0]),tags:[]};if(d&&(p.containerName=d),s.childItems&&s.childItems.length>0)for(let f of s.childItems)g(u,f,p.name);u.push(p)};let c=[];return a.forEach(u=>g(c,u)),c})}},l=class{};h(l,"unknown","");h(l,"keyword","keyword");h(l,"script","script");h(l,"module","module");h(l,"class","class");h(l,"interface","interface");h(l,"type","type");h(l,"enum","enum");h(l,"variable","var");h(l,"localVariable","local var");h(l,"function","function");h(l,"localFunction","local function");h(l,"memberFunction","method");h(l,"memberGetAccessor","getter");h(l,"memberSetAccessor","setter");h(l,"memberVariable","property");h(l,"constructorImplementation","constructor");h(l,"callSignature","call");h(l,"indexSignature","index");h(l,"constructSignature","construct");h(l,"parameter","parameter");h(l,"typeParameter","type parameter");h(l,"primitiveType","primitive type");h(l,"label","label");h(l,"alias","alias");h(l,"const","const");h(l,"let","let");h(l,"warning","warning");var _=Object.create(null);_[l.module]=n.languages.SymbolKind.Module;_[l.class]=n.languages.SymbolKind.Class;_[l.enum]=n.languages.SymbolKind.Enum;_[l.interface]=n.languages.SymbolKind.Interface;_[l.memberFunction]=n.languages.SymbolKind.Method;_[l.memberVariable]=n.languages.SymbolKind.Property;_[l.memberGetAccessor]=n.languages.SymbolKind.Property;_[l.memberSetAccessor]=n.languages.SymbolKind.Property;_[l.variable]=n.languages.SymbolKind.Variable;_[l.const]=n.languages.SymbolKind.Variable;_[l.localVariable]=n.languages.SymbolKind.Variable;_[l.variable]=n.languages.SymbolKind.Variable;_[l.function]=n.languages.SymbolKind.Function;_[l.localFunction]=n.languages.SymbolKind.Function;var v=class extends y{static _convertOptions(e){return{ConvertTabsToSpaces:e.insertSpaces,TabSize:e.tabSize,IndentSize:e.tabSize,IndentStyle:2,NewLineCharacter:`
`,InsertSpaceAfterCommaDelimiter:!0,InsertSpaceAfterSemicolonInForStatements:!0,InsertSpaceBeforeAndAfterBinaryOperators:!0,InsertSpaceAfterKeywordsInControlFlowStatements:!0,InsertSpaceAfterFunctionKeywordForAnonymousFunctions:!0,InsertSpaceAfterOpeningAndBeforeClosingNonemptyParenthesis:!1,InsertSpaceAfterOpeningAndBeforeClosingNonemptyBrackets:!1,InsertSpaceAfterOpeningAndBeforeClosingTemplateStringBraces:!1,PlaceOpenBraceOnNewLineForControlBlocks:!1,PlaceOpenBraceOnNewLineForFunctions:!1}}_convertTextChanges(e,t){return{text:t.newText,range:this._textSpanToRange(e,t.span)}}},ne=class extends v{provideDocumentRangeFormattingEdits(e,t,r,i){return m(this,null,function*(){const a=e.uri,g=e.getOffsetAt({lineNumber:t.startLineNumber,column:t.startColumn}),c=e.getOffsetAt({lineNumber:t.endLineNumber,column:t.endColumn}),u=yield this._worker(a);if(e.isDisposed())return;const s=yield u.getFormattingEditsForRange(a.toString(),g,c,v._convertOptions(r));if(!(!s||e.isDisposed()))return s.map(d=>this._convertTextChanges(e,d))})}},ae=class extends v{get autoFormatTriggerCharacters(){return[";","}",`
`]}provideOnTypeFormattingEdits(e,t,r,i,a){return m(this,null,function*(){const g=e.uri,c=e.getOffsetAt(t),u=yield this._worker(g);if(e.isDisposed())return;const s=yield u.getFormattingEditsAfterKeystroke(g.toString(),c,r,v._convertOptions(i));if(!(!s||e.isDisposed()))return s.map(d=>this._convertTextChanges(e,d))})}},oe=class extends v{provideCodeActions(e,t,r,i){return m(this,null,function*(){const a=e.uri,g=e.getOffsetAt({lineNumber:t.startLineNumber,column:t.startColumn}),c=e.getOffsetAt({lineNumber:t.endLineNumber,column:t.endColumn}),u=v._convertOptions(e.getOptions()),s=r.markers.filter(b=>b.code).map(b=>b.code).map(Number),d=yield this._worker(a);if(e.isDisposed())return;const p=yield d.getCodeFixesAtPosition(a.toString(),g,c,s,u);return!p||e.isDisposed()?{actions:[],dispose:()=>{}}:{actions:p.filter(b=>b.changes.filter(k=>k.isNewFile).length===0).map(b=>this._tsCodeFixActionToMonacoCodeAction(e,r,b)),dispose:()=>{}}})}_tsCodeFixActionToMonacoCodeAction(e,t,r){const i=[];for(const g of r.changes)for(const c of g.textChanges)i.push({resource:e.uri,edit:{range:this._textSpanToRange(e,c.span),text:c.newText}});return{title:r.description,edit:{edits:i},diagnostics:t.markers,kind:"quickfix"}}},le=class extends y{constructor(e,t){super(t),this._libFiles=e}provideRenameEdits(e,t,r,i){return m(this,null,function*(){const a=e.uri,g=a.toString(),c=e.getOffsetAt(t),u=yield this._worker(a);if(e.isDisposed())return;const s=yield u.getRenameInfo(g,c,{allowRenameOfImportPath:!1});if(s.canRename===!1)return{edits:[],rejectReason:s.localizedErrorMessage};if(s.fileToRename!==void 0)throw new Error("Renaming files is not supported.");const d=yield u.findRenameLocations(g,c,!1,!1,!1);if(!d||e.isDisposed())return;const p=[];for(const f of d){const b=this._libFiles.getOrCreateModel(f.fileName);if(b)p.push({resource:b.uri,edit:{range:this._textSpanToRange(b,f.textSpan),text:r}});else throw new Error(`Unknown file ${f.fileName}.`)}return{edits:p}})}},ce=class extends y{provideInlayHints(e,t,r){return m(this,null,function*(){const i=e.uri,a=i.toString(),g=e.getOffsetAt({lineNumber:t.startLineNumber,column:t.startColumn}),c=e.getOffsetAt({lineNumber:t.endLineNumber,column:t.endColumn}),u=yield this._worker(i);return e.isDisposed()?null:{hints:(yield u.provideInlayHints(a,g,c)).map(p=>N(I({},p),{label:p.text,position:e.getPositionAt(p.position),kind:this._convertHintKind(p.kind)})),dispose:()=>{}}})}_convertHintKind(e){switch(e){case"Parameter":return n.languages.InlayHintKind.Parameter;case"Type":return n.languages.InlayHintKind.Type;default:return n.languages.InlayHintKind.Type}}},L,P;function fe(e){P=E(e,"typescript")}function he(e){L=E(e,"javascript")}function be(){return new Promise((e,t)=>{if(!L)return t("JavaScript not registered!");e(L)})}function me(){return new Promise((e,t)=>{if(!P)return t("TypeScript not registered!");e(P)})}function E(e,t){const r=new X(t,e),i=(...g)=>r.getLanguageServiceWorker(...g),a=new Y(i);return n.languages.registerCompletionItemProvider(t,new F(i)),n.languages.registerSignatureHelpProvider(t,new R(i)),n.languages.registerHoverProvider(t,new ee(i)),n.languages.registerDocumentHighlightProvider(t,new te(i)),n.languages.registerDefinitionProvider(t,new re(a,i)),n.languages.registerReferenceProvider(t,new se(a,i)),n.languages.registerDocumentSymbolProvider(t,new ie(i)),n.languages.registerDocumentRangeFormattingEditProvider(t,new ne(i)),n.languages.registerOnTypeFormattingEditProvider(t,new ae(i)),n.languages.registerCodeActionProvider(t,new oe(i)),n.languages.registerRenameProvider(t,new le(a,i)),n.languages.registerInlayHintsProvider(t,new ce(i)),new Z(a,e,t,i),i}export{y as Adapter,oe as CodeActionAdaptor,re as DefinitionAdapter,Z as DiagnosticsAdapter,ne as FormatAdapter,v as FormatHelper,ae as FormatOnTypeAdapter,ce as InlayHintsAdapter,l as Kind,Y as LibFiles,te as OccurrencesAdapter,ie as OutlineAdapter,ee as QuickInfoAdapter,se as ReferenceAdapter,le as RenameAdapter,R as SignatureHelpAdapter,F as SuggestAdapter,X as WorkerManager,T as flattenDiagnosticMessageText,be as getJavaScriptWorker,me as getTypeScriptWorker,he as setupJavaScript,fe as setupTypeScript};
