package scopus

import (
	"strings"
	"testing"
)

func TestExtractsAbstract(t *testing.T) {
	have, err := ParseAbstract(strings.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}
	have = strings.TrimSpace(have)

	if have != want {
		t.Fatalf("want:\n%q\nhave:\n%q", want, have)
	}
}

const want = `Motion estimation (ME) is a high efficiency video coding (HEVC) process for determining motion vectors that describe the blocks transformation direction from one adjacent frame to a future frame in a video sequence. ME is a memory and computation consuming process which accounts for more than 50% of the total running time of HEVC. To conquer the memory and computation challenges, this paper presents ReME, a highly paralleled processing-in-memory (PIM) architecture for the ME process based on resistive random access memory (ReRAM). In ReME, the space of ReRAM is mainly separated into storage engine and ME processing engine. The storage engine is used as conventional memory to store video frames and intermediate data, while the computation operations of ME are performed in ME processing engines. Each ME processing engine in ReME consists of Sum of Absolute Differences (SAD) modules, interpolation modules, and Sum of Absolute Transformed Difference (SATD) modules that transfer ME functions into ReRAM-based logic analog computation units. ReME further cooperates these basic computation units to perform ME processes in a highly parallel manner. Simulation results show that the proposed ReME accelerator significantly outperforms other implementations with time consuming and energy saving. © 2021 Elsevier B.V.`

const html = `<!DOCTYPE html>
<script>
isAuthorFeedbackMVP = true
</script>
<html lang="en_US">
<head>
<META HTTP-EQUIV="CACHE-CONTROL" CONTENT="NO-CACHE">
<META HTTP-EQUIV="PRAGMA" CONTENT="NO-CACHE">
<META NAME="verify-v1" CONTENT="M5q4CxMVJyf1WE5UwVmHfV8q3LK1vh/3qmo894V0Pqg=">
<META NAME="dc.identifier" CONTENT ="10.1016/j.sysarc.2021.102123">
<script type="text/javascript">(window.NREUM||(NREUM={})).init={privacy:{cookies_enabled:true},ajax:{deny_list:["bam-cell.nr-data.net"]},distributed_tracing:{enabled:true}};(window.NREUM||(NREUM={})).loader_config={agentID:"31455944",accountID:"1281161",trustKey:"2038175",xpid:"VQQPUFdVCRADVVVXAwABVA==",licenseKey:"0268925da8",applicationID:"31454162"};window.NREUM||(NREUM={}),__nr_require=function(t,e,n){function r(n){if(!e[n]){var o=e[n]={exports:{}};t[n][0].call(o.exports,function(e){var o=t[n][1][e];return r(o||e)},o,o.exports)}return e[n].exports}if("function"==typeof __nr_require)return __nr_require;for(var o=0;o<n.length;o++)r(n[o]);return r}({1:[function(t,e,n){function r(t){try{s.console&&console.log(t)}catch(e){}}var o,i=t("ee"),a=t(32),s={};try{o=localStorage.getItem("__nr_flags").split(","),console&&"function"==typeof console.log&&(s.console=!0,o.indexOf("dev")!==-1&&(s.dev=!0),o.indexOf("nr_dev")!==-1&&(s.nrDev=!0))}catch(c){}s.nrDev&&i.on("internal-error",function(t){r(t.stack)}),s.dev&&i.on("fn-err",function(t,e,n){r(n.stack)}),s.dev&&(r("NR AGENT IN DEVELOPMENT MODE"),r("flags: "+a(s,function(t,e){return t}).join(", ")))},{}],2:[function(t,e,n){function r(t,e,n,r,s){try{l?l-=1:o(s||new UncaughtException(t,e,n),!0)}catch(f){try{i("ierr",[f,c.now(),!0])}catch(d){}}return"function"==typeof u&&u.apply(this,a(arguments))}function UncaughtException(t,e,n){this.message=t||"Uncaught error with no additional information",this.sourceURL=e,this.line=n}function o(t,e){var n=e?null:c.now();i("err",[t,n])}var i=t("handle"),a=t(33),s=t("ee"),c=t("loader"),f=t("gos"),u=window.onerror,d=!1,p="nr@seenError";if(!c.disabled){var l=0;c.features.err=!0,t(1),window.onerror=r;try{throw new Error}catch(h){"stack"in h&&(t(14),t(13),"addEventListener"in window&&t(7),c.xhrWrappable&&t(15),d=!0)}s.on("fn-start",function(t,e,n){d&&(l+=1)}),s.on("fn-err",function(t,e,n){d&&!n[p]&&(f(n,p,function(){return!0}),this.thrown=!0,o(n))}),s.on("fn-end",function(){d&&!this.thrown&&l>0&&(l-=1)}),s.on("internal-error",function(t){i("ierr",[t,c.now(),!0])})}},{}],3:[function(t,e,n){var r=t("loader");r.disabled||(r.features.ins=!0)},{}],4:[function(t,e,n){function r(){U++,L=g.hash,this[u]=y.now()}function o(){U--,g.hash!==L&&i(0,!0);var t=y.now();this[h]=~~this[h]+t-this[u],this[d]=t}function i(t,e){E.emit("newURL",[""+g,e])}function a(t,e){t.on(e,function(){this[e]=y.now()})}var s="-start",c="-end",f="-body",u="fn"+s,d="fn"+c,p="cb"+s,l="cb"+c,h="jsTime",m="fetch",v="addEventListener",w=window,g=w.location,y=t("loader");if(w[v]&&y.xhrWrappable&&!y.disabled){var x=t(11),b=t(12),E=t(9),R=t(7),O=t(14),T=t(8),P=t(15),S=t(10),M=t("ee"),N=M.get("tracer"),C=t(23);t(17),y.features.spa=!0;var L,U=0;M.on(u,r),b.on(p,r),S.on(p,r),M.on(d,o),b.on(l,o),S.on(l,o),M.buffer([u,d,"xhr-resolved"]),R.buffer([u]),O.buffer(["setTimeout"+c,"clearTimeout"+s,u]),P.buffer([u,"new-xhr","send-xhr"+s]),T.buffer([m+s,m+"-done",m+f+s,m+f+c]),E.buffer(["newURL"]),x.buffer([u]),b.buffer(["propagate",p,l,"executor-err","resolve"+s]),N.buffer([u,"no-"+u]),S.buffer(["new-jsonp","cb-start","jsonp-error","jsonp-end"]),a(T,m+s),a(T,m+"-done"),a(S,"new-jsonp"),a(S,"jsonp-end"),a(S,"cb-start"),E.on("pushState-end",i),E.on("replaceState-end",i),w[v]("hashchange",i,C(!0)),w[v]("load",i,C(!0)),w[v]("popstate",function(){i(0,U>1)},C(!0))}},{}],5:[function(t,e,n){function r(){var t=new PerformanceObserver(function(t,e){var n=t.getEntries();s(v,[n])});try{t.observe({entryTypes:["resource"]})}catch(e){}}function o(t){if(s(v,[window.performance.getEntriesByType(w)]),window.performance["c"+p])try{window.performance[h](m,o,!1)}catch(t){}else try{window.performance[h]("webkit"+m,o,!1)}catch(t){}}function i(t){}if(window.performance&&window.performance.timing&&window.performance.getEntriesByType){var a=t("ee"),s=t("handle"),c=t(14),f=t(13),u=t(6),d=t(23),p="learResourceTimings",l="addEventListener",h="removeEventListener",m="resourcetimingbufferfull",v="bstResource",w="resource",g="-start",y="-end",x="fn"+g,b="fn"+y,E="bstTimer",R="pushState",O=t("loader");if(!O.disabled){O.features.stn=!0,t(9),"addEventListener"in window&&t(7);var T=NREUM.o.EV;a.on(x,function(t,e){var n=t[0];n instanceof T&&(this.bstStart=O.now())}),a.on(b,function(t,e){var n=t[0];n instanceof T&&s("bst",[n,e,this.bstStart,O.now()])}),c.on(x,function(t,e,n){this.bstStart=O.now(),this.bstType=n}),c.on(b,function(t,e){s(E,[e,this.bstStart,O.now(),this.bstType])}),f.on(x,function(){this.bstStart=O.now()}),f.on(b,function(t,e){s(E,[e,this.bstStart,O.now(),"requestAnimationFrame"])}),a.on(R+g,function(t){this.time=O.now(),this.startPath=location.pathname+location.hash}),a.on(R+y,function(t){s("bstHist",[location.pathname+location.hash,this.startPath,this.time])}),u()?(s(v,[window.performance.getEntriesByType("resource")]),r()):l in window.performance&&(window.performance["c"+p]?window.performance[l](m,o,d(!1)):window.performance[l]("webkit"+m,o,d(!1))),document[l]("scroll",i,d(!1)),document[l]("keypress",i,d(!1)),document[l]("click",i,d(!1))}}},{}],6:[function(t,e,n){e.exports=function(){return"PerformanceObserver"in window&&"function"==typeof window.PerformanceObserver}},{}],7:[function(t,e,n){function r(t){for(var e=t;e&&!e.hasOwnProperty(u);)e=Object.getPrototypeOf(e);e&&o(e)}function o(t){s.inPlace(t,[u,d],"-",i)}function i(t,e){return t[1]}var a=t("ee").get("events"),s=t("wrap-function")(a,!0),c=t("gos"),f=XMLHttpRequest,u="addEventListener",d="removeEventListener";e.exports=a,"getPrototypeOf"in Object?(r(document),r(window),r(f.prototype)):f.prototype.hasOwnProperty(u)&&(o(window),o(f.prototype)),a.on(u+"-start",function(t,e){var n=t[1];if(null!==n&&("function"==typeof n||"object"==typeof n)){var r=c(n,"nr@wrapped",function(){function t(){if("function"==typeof n.handleEvent)return n.handleEvent.apply(n,arguments)}var e={object:t,"function":n}[typeof n];return e?s(e,"fn-",null,e.name||"anonymous"):n});this.wrapped=t[1]=r}}),a.on(d+"-start",function(t){t[1]=this.wrapped||t[1]})},{}],8:[function(t,e,n){function r(t,e,n){var r=t[e];"function"==typeof r&&(t[e]=function(){var t=i(arguments),e={};o.emit(n+"before-start",[t],e);var a;e[m]&&e[m].dt&&(a=e[m].dt);var s=r.apply(this,t);return o.emit(n+"start",[t,a],s),s.then(function(t){return o.emit(n+"end",[null,t],s),t},function(t){throw o.emit(n+"end",[t],s),t})})}var o=t("ee").get("fetch"),i=t(33),a=t(32);e.exports=o;var s=window,c="fetch-",f=c+"body-",u=["arrayBuffer","blob","json","text","formData"],d=s.Request,p=s.Response,l=s.fetch,h="prototype",m="nr@context";d&&p&&l&&(a(u,function(t,e){r(d[h],e,f),r(p[h],e,f)}),r(s,"fetch",c),o.on(c+"end",function(t,e){var n=this;if(e){var r=e.headers.get("content-length");null!==r&&(n.rxSize=r),o.emit(c+"done",[null,e],n)}else o.emit(c+"done",[t],n)}))},{}],9:[function(t,e,n){var r=t("ee").get("history"),o=t("wrap-function")(r);e.exports=r;var i=window.history&&window.history.constructor&&window.history.constructor.prototype,a=window.history;i&&i.pushState&&i.replaceState&&(a=i),o.inPlace(a,["pushState","replaceState"],"-")},{}],10:[function(t,e,n){function r(t){function e(){f.emit("jsonp-end",[],l),t.removeEventListener("load",e,c(!1)),t.removeEventListener("error",n,c(!1))}function n(){f.emit("jsonp-error",[],l),f.emit("jsonp-end",[],l),t.removeEventListener("load",e,c(!1)),t.removeEventListener("error",n,c(!1))}var r=t&&"string"==typeof t.nodeName&&"script"===t.nodeName.toLowerCase();if(r){var o="function"==typeof t.addEventListener;if(o){var a=i(t.src);if(a){var d=s(a),p="function"==typeof d.parent[d.key];if(p){var l={};u.inPlace(d.parent,[d.key],"cb-",l),t.addEventListener("load",e,c(!1)),t.addEventListener("error",n,c(!1)),f.emit("new-jsonp",[t.src],l)}}}}}function o(){return"addEventListener"in window}function i(t){var e=t.match(d);return e?e[1]:null}function a(t,e){var n=t.match(l),r=n[1],o=n[3];return o?a(o,e[r]):e[r]}function s(t){var e=t.match(p);return e&&e.length>=3?{key:e[2],parent:a(e[1],window)}:{key:t,parent:window}}var c=t(23),f=t("ee").get("jsonp"),u=t("wrap-function")(f);if(e.exports=f,o()){var d=/[?&](?:callback|cb)=([^&#]+)/,p=/(.*)\.([^.]+)/,l=/^(\w+)(\.|$)(.*)$/,h=["appendChild","insertBefore","replaceChild"];Node&&Node.prototype&&Node.prototype.appendChild?u.inPlace(Node.prototype,h,"dom-"):(u.inPlace(HTMLElement.prototype,h,"dom-"),u.inPlace(HTMLHeadElement.prototype,h,"dom-"),u.inPlace(HTMLBodyElement.prototype,h,"dom-")),f.on("dom-start",function(t){r(t[0])})}},{}],11:[function(t,e,n){var r=t("ee").get("mutation"),o=t("wrap-function")(r),i=NREUM.o.MO;e.exports=r,i&&(window.MutationObserver=function(t){return this instanceof i?new i(o(t,"fn-")):i.apply(this,arguments)},MutationObserver.prototype=i.prototype)},{}],12:[function(t,e,n){function r(t){var e=i.context(),n=s(t,"executor-",e,null,!1),r=new f(n);return i.context(r).getCtx=function(){return e},r}var o=t("wrap-function"),i=t("ee").get("promise"),a=t("ee").getOrSetContext,s=o(i),c=t(32),f=NREUM.o.PR;e.exports=i,f&&(window.Promise=r,["all","race"].forEach(function(t){var e=f[t];f[t]=function(n){function r(t){return function(){i.emit("propagate",[null,!o],a,!1,!1),o=o||!t}}var o=!1;c(n,function(e,n){Promise.resolve(n).then(r("all"===t),r(!1))});var a=e.apply(f,arguments),s=f.resolve(a);return s}}),["resolve","reject"].forEach(function(t){var e=f[t];f[t]=function(t){var n=e.apply(f,arguments);return t!==n&&i.emit("propagate",[t,!0],n,!1,!1),n}}),f.prototype["catch"]=function(t){return this.then(null,t)},f.prototype=Object.create(f.prototype,{constructor:{value:r}}),c(Object.getOwnPropertyNames(f),function(t,e){try{r[e]=f[e]}catch(n){}}),o.wrapInPlace(f.prototype,"then",function(t){return function(){var e=this,n=o.argsToArray.apply(this,arguments),r=a(e);r.promise=e,n[0]=s(n[0],"cb-",r,null,!1),n[1]=s(n[1],"cb-",r,null,!1);var c=t.apply(this,n);return r.nextPromise=c,i.emit("propagate",[e,!0],c,!1,!1),c}}),i.on("executor-start",function(t){t[0]=s(t[0],"resolve-",this,null,!1),t[1]=s(t[1],"resolve-",this,null,!1)}),i.on("executor-err",function(t,e,n){t[1](n)}),i.on("cb-end",function(t,e,n){i.emit("propagate",[n,!0],this.nextPromise,!1,!1)}),i.on("propagate",function(t,e,n){this.getCtx&&!e||(this.getCtx=function(){if(t instanceof Promise)var e=i.context(t);return e&&e.getCtx?e.getCtx():this})}),r.toString=function(){return""+f})},{}],13:[function(t,e,n){var r=t("ee").get("raf"),o=t("wrap-function")(r),i="equestAnimationFrame";e.exports=r,o.inPlace(window,["r"+i,"mozR"+i,"webkitR"+i,"msR"+i],"raf-"),r.on("raf-start",function(t){t[0]=o(t[0],"fn-")})},{}],14:[function(t,e,n){function r(t,e,n){t[0]=a(t[0],"fn-",null,n)}function o(t,e,n){this.method=n,this.timerDuration=isNaN(t[1])?0:+t[1],t[0]=a(t[0],"fn-",this,n)}var i=t("ee").get("timer"),a=t("wrap-function")(i),s="setTimeout",c="setInterval",f="clearTimeout",u="-start",d="-";e.exports=i,a.inPlace(window,[s,"setImmediate"],s+d),a.inPlace(window,[c],c+d),a.inPlace(window,[f,"clearImmediate"],f+d),i.on(c+u,r),i.on(s+u,o)},{}],15:[function(t,e,n){function r(t,e){d.inPlace(e,["onreadystatechange"],"fn-",s)}function o(){var t=this,e=u.context(t);t.readyState>3&&!e.resolved&&(e.resolved=!0,u.emit("xhr-resolved",[],t)),d.inPlace(t,y,"fn-",s)}function i(t){x.push(t),m&&(E?E.then(a):w?w(a):(R=-R,O.data=R))}function a(){for(var t=0;t<x.length;t++)r([],x[t]);x.length&&(x=[])}function s(t,e){return e}function c(t,e){for(var n in t)e[n]=t[n];return e}t(7);var f=t("ee"),u=f.get("xhr"),d=t("wrap-function")(u),p=t(23),l=NREUM.o,h=l.XHR,m=l.MO,v=l.PR,w=l.SI,g="readystatechange",y=["onload","onerror","onabort","onloadstart","onloadend","onprogress","ontimeout"],x=[];e.exports=u;var b=window.XMLHttpRequest=function(t){var e=new h(t);try{u.emit("new-xhr",[e],e),e.addEventListener(g,o,p(!1))}catch(n){try{u.emit("internal-error",[n])}catch(r){}}return e};if(c(h,b),b.prototype=h.prototype,d.inPlace(b.prototype,["open","send"],"-xhr-",s),u.on("send-xhr-start",function(t,e){r(t,e),i(e)}),u.on("open-xhr-start",r),m){var E=v&&v.resolve();if(!w&&!v){var R=1,O=document.createTextNode(R);new m(a).observe(O,{characterData:!0})}}else f.on("fn-end",function(t){t[0]&&t[0].type===g||a()})},{}],16:[function(t,e,n){function r(t){if(!s(t))return null;var e=window.NREUM;if(!e.loader_config)return null;var n=(e.loader_config.accountID||"").toString()||null,r=(e.loader_config.agentID||"").toString()||null,f=(e.loader_config.trustKey||"").toString()||null;if(!n||!r)return null;var h=l.generateSpanId(),m=l.generateTraceId(),v=Date.now(),w={spanId:h,traceId:m,timestamp:v};return(t.sameOrigin||c(t)&&p())&&(w.traceContextParentHeader=o(h,m),w.traceContextStateHeader=i(h,v,n,r,f)),(t.sameOrigin&&!u()||!t.sameOrigin&&c(t)&&d())&&(w.newrelicHeader=a(h,m,v,n,r,f)),w}function o(t,e){return"00-"+e+"-"+t+"-01"}function i(t,e,n,r,o){var i=0,a="",s=1,c="",f="";return o+"@nr="+i+"-"+s+"-"+n+"-"+r+"-"+t+"-"+a+"-"+c+"-"+f+"-"+e}function a(t,e,n,r,o,i){var a="btoa"in window&&"function"==typeof window.btoa;if(!a)return null;var s={v:[0,1],d:{ty:"Browser",ac:r,ap:o,id:t,tr:e,ti:n}};return i&&r!==i&&(s.d.tk=i),btoa(JSON.stringify(s))}function s(t){return f()&&c(t)}function c(t){var e=!1,n={};if("init"in NREUM&&"distributed_tracing"in NREUM.init&&(n=NREUM.init.distributed_tracing),t.sameOrigin)e=!0;else if(n.allowed_origins instanceof Array)for(var r=0;r<n.allowed_origins.length;r++){var o=h(n.allowed_origins[r]);if(t.hostname===o.hostname&&t.protocol===o.protocol&&t.port===o.port){e=!0;break}}return e}function f(){return"init"in NREUM&&"distributed_tracing"in NREUM.init&&!!NREUM.init.distributed_tracing.enabled}function u(){return"init"in NREUM&&"distributed_tracing"in NREUM.init&&!!NREUM.init.distributed_tracing.exclude_newrelic_header}function d(){return"init"in NREUM&&"distributed_tracing"in NREUM.init&&NREUM.init.distributed_tracing.cors_use_newrelic_header!==!1}function p(){return"init"in NREUM&&"distributed_tracing"in NREUM.init&&!!NREUM.init.distributed_tracing.cors_use_tracecontext_headers}var l=t(29),h=t(18);e.exports={generateTracePayload:r,shouldGenerateTrace:s}},{}],17:[function(t,e,n){function r(t){var e=this.params,n=this.metrics;if(!this.ended){this.ended=!0;for(var r=0;r<p;r++)t.removeEventListener(d[r],this.listener,!1);e.aborted||(n.duration=a.now()-this.startTime,this.loadCaptureCalled||4!==t.readyState?null==e.status&&(e.status=0):i(this,t),n.cbTime=this.cbTime,s("xhr",[e,n,this.startTime,this.endTime,"xhr"],this))}}function o(t,e){var n=c(e),r=t.params;r.hostname=n.hostname,r.port=n.port,r.protocol=n.protocol,r.host=n.hostname+":"+n.port,r.pathname=n.pathname,t.parsedOrigin=n,t.sameOrigin=n.sameOrigin}function i(t,e){t.params.status=e.status;var n=v(e,t.lastSize);if(n&&(t.metrics.rxSize=n),t.sameOrigin){var r=e.getResponseHeader("X-NewRelic-App-Data");r&&(t.params.cat=r.split(", ").pop())}t.loadCaptureCalled=!0}var a=t("loader");if(a.xhrWrappable&&!a.disabled){var s=t("handle"),c=t(18),f=t(16).generateTracePayload,u=t("ee"),d=["load","error","abort","timeout"],p=d.length,l=t("id"),h=t(24),m=t(22),v=t(19),w=t(23),g=NREUM.o.REQ,y=window.XMLHttpRequest;a.features.xhr=!0,t(15),t(8),u.on("new-xhr",function(t){var e=this;e.totalCbs=0,e.called=0,e.cbTime=0,e.end=r,e.ended=!1,e.xhrGuids={},e.lastSize=null,e.loadCaptureCalled=!1,e.params=this.params||{},e.metrics=this.metrics||{},t.addEventListener("load",function(n){i(e,t)},w(!1)),h&&(h>34||h<10)||t.addEventListener("progress",function(t){e.lastSize=t.loaded},w(!1))}),u.on("open-xhr-start",function(t){this.params={method:t[0]},o(this,t[1]),this.metrics={}}),u.on("open-xhr-end",function(t,e){"loader_config"in NREUM&&"xpid"in NREUM.loader_config&&this.sameOrigin&&e.setRequestHeader("X-NewRelic-ID",NREUM.loader_config.xpid);var n=f(this.parsedOrigin);if(n){var r=!1;n.newrelicHeader&&(e.setRequestHeader("newrelic",n.newrelicHeader),r=!0),n.traceContextParentHeader&&(e.setRequestHeader("traceparent",n.traceContextParentHeader),n.traceContextStateHeader&&e.setRequestHeader("tracestate",n.traceContextStateHeader),r=!0),r&&(this.dt=n)}}),u.on("send-xhr-start",function(t,e){var n=this.metrics,r=t[0],o=this;if(n&&r){var i=m(r);i&&(n.txSize=i)}this.startTime=a.now(),this.listener=function(t){try{"abort"!==t.type||o.loadCaptureCalled||(o.params.aborted=!0),("load"!==t.type||o.called===o.totalCbs&&(o.onloadCalled||"function"!=typeof e.onload))&&o.end(e)}catch(n){try{u.emit("internal-error",[n])}catch(r){}}};for(var s=0;s<p;s++)e.addEventListener(d[s],this.listener,w(!1))}),u.on("xhr-cb-time",function(t,e,n){this.cbTime+=t,e?this.onloadCalled=!0:this.called+=1,this.called!==this.totalCbs||!this.onloadCalled&&"function"==typeof n.onload||this.end(n)}),u.on("xhr-load-added",function(t,e){var n=""+l(t)+!!e;this.xhrGuids&&!this.xhrGuids[n]&&(this.xhrGuids[n]=!0,this.totalCbs+=1)}),u.on("xhr-load-removed",function(t,e){var n=""+l(t)+!!e;this.xhrGuids&&this.xhrGuids[n]&&(delete this.xhrGuids[n],this.totalCbs-=1)}),u.on("xhr-resolved",function(){this.endTime=a.now()}),u.on("addEventListener-end",function(t,e){e instanceof y&&"load"===t[0]&&u.emit("xhr-load-added",[t[1],t[2]],e)}),u.on("removeEventListener-end",function(t,e){e instanceof y&&"load"===t[0]&&u.emit("xhr-load-removed",[t[1],t[2]],e)}),u.on("fn-start",function(t,e,n){e instanceof y&&("onload"===n&&(this.onload=!0),("load"===(t[0]&&t[0].type)||this.onload)&&(this.xhrCbStart=a.now()))}),u.on("fn-end",function(t,e){this.xhrCbStart&&u.emit("xhr-cb-time",[a.now()-this.xhrCbStart,this.onload,e],e)}),u.on("fetch-before-start",function(t){function e(t,e){var n=!1;return e.newrelicHeader&&(t.set("newrelic",e.newrelicHeader),n=!0),e.traceContextParentHeader&&(t.set("traceparent",e.traceContextParentHeader),e.traceContextStateHeader&&t.set("tracestate",e.traceContextStateHeader),n=!0),n}var n,r=t[1]||{};"string"==typeof t[0]?n=t[0]:t[0]&&t[0].url?n=t[0].url:window.URL&&t[0]&&t[0]instanceof URL&&(n=t[0].href),n&&(this.parsedOrigin=c(n),this.sameOrigin=this.parsedOrigin.sameOrigin);var o=f(this.parsedOrigin);if(o&&(o.newrelicHeader||o.traceContextParentHeader))if("string"==typeof t[0]||window.URL&&t[0]&&t[0]instanceof URL){var i={};for(var a in r)i[a]=r[a];i.headers=new Headers(r.headers||{}),e(i.headers,o)&&(this.dt=o),t.length>1?t[1]=i:t.push(i)}else t[0]&&t[0].headers&&e(t[0].headers,o)&&(this.dt=o)}),u.on("fetch-start",function(t,e){this.params={},this.metrics={},this.startTime=a.now(),this.dt=e,t.length>=1&&(this.target=t[0]),t.length>=2&&(this.opts=t[1]);var n,r=this.opts||{},i=this.target;"string"==typeof i?n=i:"object"==typeof i&&i instanceof g?n=i.url:window.URL&&"object"==typeof i&&i instanceof URL&&(n=i.href),o(this,n);var s=(""+(i&&i instanceof g&&i.method||r.method||"GET")).toUpperCase();this.params.method=s,this.txSize=m(r.body)||0}),u.on("fetch-done",function(t,e){this.endTime=a.now(),this.params||(this.params={}),this.params.status=e?e.status:0;var n;"string"==typeof this.rxSize&&this.rxSize.length>0&&(n=+this.rxSize);var r={txSize:this.txSize,rxSize:n,duration:a.now()-this.startTime};s("xhr",[this.params,r,this.startTime,this.endTime,"fetch"],this)})}},{}],18:[function(t,e,n){var r={};e.exports=function(t){if(t in r)return r[t];var e=document.createElement("a"),n=window.location,o={};e.href=t,o.port=e.port;var i=e.href.split("://");!o.port&&i[1]&&(o.port=i[1].split("/")[0].split("@").pop().split(":")[1]),o.port&&"0"!==o.port||(o.port="https"===i[0]?"443":"80"),o.hostname=e.hostname||n.hostname,o.pathname=e.pathname,o.protocol=i[0],"/"!==o.pathname.charAt(0)&&(o.pathname="/"+o.pathname);var a=!e.protocol||":"===e.protocol||e.protocol===n.protocol,s=e.hostname===document.domain&&e.port===n.port;return o.sameOrigin=a&&(!e.hostname||s),"/"===o.pathname&&(r[t]=o),o}},{}],19:[function(t,e,n){function r(t,e){var n=t.responseType;return"json"===n&&null!==e?e:"arraybuffer"===n||"blob"===n||"json"===n?o(t.response):"text"===n||""===n||void 0===n?o(t.responseText):void 0}var o=t(22);e.exports=r},{}],20:[function(t,e,n){function r(){}function o(t,e,n,r){return function(){return u.recordSupportability("API/"+e+"/called"),i(t+e,[f.now()].concat(s(arguments)),n?null:this,r),n?void 0:this}}var i=t("handle"),a=t(32),s=t(33),c=t("ee").get("tracer"),f=t("loader"),u=t(25),d=NREUM;"undefined"==typeof window.newrelic&&(newrelic=d);var p=["setPageViewName","setCustomAttribute","setErrorHandler","finished","addToTrace","inlineHit","addRelease"],l="api-",h=l+"ixn-";a(p,function(t,e){d[e]=o(l,e,!0,"api")}),d.addPageAction=o(l,"addPageAction",!0),d.setCurrentRouteName=o(l,"routeName",!0),e.exports=newrelic,d.interaction=function(){return(new r).get()};var m=r.prototype={createTracer:function(t,e){var n={},r=this,o="function"==typeof e;return i(h+"tracer",[f.now(),t,n],r),function(){if(c.emit((o?"":"no-")+"fn-start",[f.now(),r,o],n),o)try{return e.apply(this,arguments)}catch(t){throw c.emit("fn-err",[arguments,this,t],n),t}finally{c.emit("fn-end",[f.now()],n)}}}};a("actionText,setName,setAttribute,save,ignore,onEnd,getContext,end,get".split(","),function(t,e){m[e]=o(h,e)}),newrelic.noticeError=function(t,e){"string"==typeof t&&(t=new Error(t)),u.recordSupportability("API/noticeError/called"),i("err",[t,f.now(),!1,e])}},{}],21:[function(t,e,n){function r(t){if(NREUM.init){for(var e=NREUM.init,n=t.split("."),r=0;r<n.length-1;r++)if(e=e[n[r]],"object"!=typeof e)return;return e=e[n[n.length-1]]}}e.exports={getConfiguration:r}},{}],22:[function(t,e,n){e.exports=function(t){if("string"==typeof t&&t.length)return t.length;if("object"==typeof t){if("undefined"!=typeof ArrayBuffer&&t instanceof ArrayBuffer&&t.byteLength)return t.byteLength;if("undefined"!=typeof Blob&&t instanceof Blob&&t.size)return t.size;if(!("undefined"!=typeof FormData&&t instanceof FormData))try{return JSON.stringify(t).length}catch(e){return}}}},{}],23:[function(t,e,n){var r=!1;try{var o=Object.defineProperty({},"passive",{get:function(){r=!0}});window.addEventListener("testPassive",null,o),window.removeEventListener("testPassive",null,o)}catch(i){}e.exports=function(t){return r?{passive:!0,capture:!!t}:!!t}},{}],24:[function(t,e,n){var r=0,o=navigator.userAgent.match(/Firefox[\/\s](\d+\.\d+)/);o&&(r=+o[1]),e.exports=r},{}],25:[function(t,e,n){function r(t,e){var n=[a,t,{name:t},e];return i("storeMetric",n,null,"api"),n}function o(t,e){var n=[s,t,{name:t},e];return i("storeEventMetrics",n,null,"api"),n}var i=t("handle"),a="sm",s="cm";e.exports={constants:{SUPPORTABILITY_METRIC:a,CUSTOM_METRIC:s},recordSupportability:r,recordCustom:o}},{}],26:[function(t,e,n){function r(){return s.exists&&performance.now?Math.round(performance.now()):(i=Math.max((new Date).getTime(),i))-a}function o(){return i}var i=(new Date).getTime(),a=i,s=t(34);e.exports=r,e.exports.offset=a,e.exports.getLastTimestamp=o},{}],27:[function(t,e,n){function r(t){return!(!t||!t.protocol||"file:"===t.protocol)}e.exports=r},{}],28:[function(t,e,n){function r(t,e){var n=t.getEntries();n.forEach(function(t){"first-paint"===t.name?p("timing",["fp",Math.floor(t.startTime)]):"first-contentful-paint"===t.name&&p("timing",["fcp",Math.floor(t.startTime)])})}function o(t,e){var n=t.getEntries();if(n.length>0){var r=n[n.length-1];if(c&&c<r.startTime)return;p("lcp",[r])}}function i(t){t.getEntries().forEach(function(t){t.hadRecentInput||p("cls",[t])})}function a(t){if(t instanceof v&&!g){var e=Math.round(t.timeStamp),n={type:t.type};e<=l.now()?n.fid=l.now()-e:e>l.offset&&e<=Date.now()?(e-=l.offset,n.fid=l.now()-e):e=l.now(),g=!0,p("timing",["fi",e,n])}}function s(t){"hidden"===t&&(c=l.now(),p("pageHide",[c]))}if(!("init"in NREUM&&"page_view_timing"in NREUM.init&&"enabled"in NREUM.init.page_view_timing&&NREUM.init.page_view_timing.enabled===!1)){var c,f,u,d,p=t("handle"),l=t("loader"),h=t(31),m=t(23),v=NREUM.o.EV;if("PerformanceObserver"in window&&"function"==typeof window.PerformanceObserver){f=new PerformanceObserver(r);try{f.observe({entryTypes:["paint"]})}catch(w){}u=new PerformanceObserver(o);try{u.observe({entryTypes:["largest-contentful-paint"]})}catch(w){}d=new PerformanceObserver(i);try{d.observe({type:"layout-shift",buffered:!0})}catch(w){}}if("addEventListener"in document){var g=!1,y=["click","keydown","mousedown","pointerdown","touchstart"];y.forEach(function(t){document.addEventListener(t,a,m(!1))})}h(s)}},{}],29:[function(t,e,n){function r(){function t(){return e?15&e[n++]:16*Math.random()|0}var e=null,n=0,r=window.crypto||window.msCrypto;r&&r.getRandomValues&&(e=r.getRandomValues(new Uint8Array(31)));for(var o,i="xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx",a="",s=0;s<i.length;s++)o=i[s],"x"===o?a+=t().toString(16):"y"===o?(o=3&t()|8,a+=o.toString(16)):a+=o;return a}function o(){return a(16)}function i(){return a(32)}function a(t){function e(){return n?15&n[r++]:16*Math.random()|0}var n=null,r=0,o=window.crypto||window.msCrypto;o&&o.getRandomValues&&Uint8Array&&(n=o.getRandomValues(new Uint8Array(31)));for(var i=[],a=0;a<t;a++)i.push(e().toString(16));return i.join("")}e.exports={generateUuid:r,generateSpanId:o,generateTraceId:i}},{}],30:[function(t,e,n){function r(t,e){if(!o)return!1;if(t!==o)return!1;if(!e)return!0;if(!i)return!1;for(var n=i.split("."),r=e.split("."),a=0;a<r.length;a++)if(r[a]!==n[a])return!1;return!0}var o=null,i=null,a=/Version\/(\S+)\s+Safari/;if(navigator.userAgent){var s=navigator.userAgent,c=s.match(a);c&&s.indexOf("Chrome")===-1&&s.indexOf("Chromium")===-1&&(o="Safari",i=c[1])}e.exports={agent:o,version:i,match:r}},{}],31:[function(t,e,n){function r(t){function e(){t(s&&document[s]?document[s]:document[i]?"hidden":"visible")}"addEventListener"in document&&a&&document.addEventListener(a,e,o(!1))}var o=t(23);e.exports=r;var i,a,s;"undefined"!=typeof document.hidden?(i="hidden",a="visibilitychange",s="visibilityState"):"undefined"!=typeof document.msHidden?(i="msHidden",a="msvisibilitychange"):"undefined"!=typeof document.webkitHidden&&(i="webkitHidden",a="webkitvisibilitychange",s="webkitVisibilityState")},{}],32:[function(t,e,n){function r(t,e){var n=[],r="",i=0;for(r in t)o.call(t,r)&&(n[i]=e(r,t[r]),i+=1);return n}var o=Object.prototype.hasOwnProperty;e.exports=r},{}],33:[function(t,e,n){function r(t,e,n){e||(e=0),"undefined"==typeof n&&(n=t?t.length:0);for(var r=-1,o=n-e||0,i=Array(o<0?0:o);++r<o;)i[r]=t[e+r];return i}e.exports=r},{}],34:[function(t,e,n){e.exports={exists:"undefined"!=typeof window.performance&&window.performance.timing&&"undefined"!=typeof window.performance.timing.navigationStart}},{}],ee:[function(t,e,n){function r(){}function o(t){function e(t){return t&&t instanceof r?t:t?f(t,c,a):a()}function n(n,r,o,i,a){if(a!==!1&&(a=!0),!l.aborted||i){t&&a&&t(n,r,o);for(var s=e(o),c=m(n),f=c.length,u=0;u<f;u++)c[u].apply(s,r);var p=d[y[n]];return p&&p.push([x,n,r,s]),s}}function i(t,e){g[t]=m(t).concat(e)}function h(t,e){var n=g[t];if(n)for(var r=0;r<n.length;r++)n[r]===e&&n.splice(r,1)}function m(t){return g[t]||[]}function v(t){return p[t]=p[t]||o(n)}function w(t,e){l.aborted||u(t,function(t,n){e=e||"feature",y[n]=e,e in d||(d[e]=[])})}var g={},y={},x={on:i,addEventListener:i,removeEventListener:h,emit:n,get:v,listeners:m,context:e,buffer:w,abort:s,aborted:!1};return x}function i(t){return f(t,c,a)}function a(){return new r}function s(){(d.api||d.feature)&&(l.aborted=!0,d=l.backlog={})}var c="nr@context",f=t("gos"),u=t(32),d={},p={},l=e.exports=o();e.exports.getOrSetContext=i,l.backlog=d},{}],gos:[function(t,e,n){function r(t,e,n){if(o.call(t,e))return t[e];var r=n();if(Object.defineProperty&&Object.keys)try{return Object.defineProperty(t,e,{value:r,writable:!0,enumerable:!1}),r}catch(i){}return t[e]=r,r}var o=Object.prototype.hasOwnProperty;e.exports=r},{}],handle:[function(t,e,n){function r(t,e,n,r){o.buffer([t],r),o.emit(t,e,n)}var o=t("ee").get("handle");e.exports=r,r.ee=o},{}],id:[function(t,e,n){function r(t){var e=typeof t;return!t||"object"!==e&&"function"!==e?-1:t===window?0:a(t,i,function(){return o++})}var o=1,i="nr@id",a=t("gos");e.exports=r},{}],loader:[function(t,e,n){function r(){if(!P++){var t=T.info=NREUM.info,e=v.getElementsByTagName("script")[0];if(setTimeout(f.abort,3e4),!(t&&t.licenseKey&&t.applicationID&&e))return f.abort();c(R,function(e,n){t[e]||(t[e]=n)});var n=a();s("mark",["onload",n+T.offset],null,"api"),s("timing",["load",n]);var r=v.createElement("script");0===t.agent.indexOf("http://")||0===t.agent.indexOf("https://")?r.src=t.agent:r.src=h+"://"+t.agent,e.parentNode.insertBefore(r,e)}}function o(){"complete"===v.readyState&&i()}function i(){s("mark",["domContent",a()+T.offset],null,"api")}var a=t(26),s=t("handle"),c=t(32),f=t("ee"),u=t(30),d=t(27),p=t(21),l=t(23),h=p.getConfiguration("ssl")===!1?"http":"https",m=window,v=m.document,w="addEventListener",g="attachEvent",y=m.XMLHttpRequest,x=y&&y.prototype,b=!d(m.location);NREUM.o={ST:setTimeout,SI:m.setImmediate,CT:clearTimeout,XHR:y,REQ:m.Request,EV:m.Event,PR:m.Promise,MO:m.MutationObserver};var E=""+location,R={beacon:"bam.nr-data.net",errorBeacon:"bam.nr-data.net",agent:"js-agent.newrelic.com/nr-spa-1212.min.js"},O=y&&x&&x[w]&&!/CriOS/.test(navigator.userAgent),T=e.exports={offset:a.getLastTimestamp(),now:a,origin:E,features:{},xhrWrappable:O,userAgent:u,disabled:b};if(!b){t(20),t(28),v[w]?(v[w]("DOMContentLoaded",i,l(!1)),m[w]("load",r,l(!1))):(v[g]("onreadystatechange",o),m[g]("onload",r)),s("mark",["firstbyte",a.getLastTimestamp()],null,"api");var P=0}},{}],"wrap-function":[function(t,e,n){function r(t,e){function n(e,n,r,c,f){function nrWrapper(){var i,a,u,p;try{a=this,i=d(arguments),u="function"==typeof r?r(i,a):r||{}}catch(l){o([l,"",[i,a,c],u],t)}s(n+"start",[i,a,c],u,f);try{return p=e.apply(a,i)}catch(h){throw s(n+"err",[i,a,h],u,f),h}finally{s(n+"end",[i,a,p],u,f)}}return a(e)?e:(n||(n=""),nrWrapper[p]=e,i(e,nrWrapper,t),nrWrapper)}function r(t,e,r,o,i){r||(r="");var s,c,f,u="-"===r.charAt(0);for(f=0;f<e.length;f++)c=e[f],s=t[c],a(s)||(t[c]=n(s,u?c+r:r,o,c,i))}function s(n,r,i,a){if(!h||e){var s=h;h=!0;try{t.emit(n,r,i,e,a)}catch(c){o([c,n,r,i],t)}h=s}}return t||(t=u),n.inPlace=r,n.flag=p,n}function o(t,e){e||(e=u);try{e.emit("internal-error",t)}catch(n){}}function i(t,e,n){if(Object.defineProperty&&Object.keys)try{var r=Object.keys(t);return r.forEach(function(n){Object.defineProperty(e,n,{get:function(){return t[n]},set:function(e){return t[n]=e,e}})}),e}catch(i){o([i],n)}for(var a in t)l.call(t,a)&&(e[a]=t[a]);return e}function a(t){return!(t&&t instanceof Function&&t.apply&&!t[p])}function s(t,e){var n=e(t);return n[p]=t,i(t,n,u),n}function c(t,e,n){var r=t[e];t[e]=s(r,n)}function f(){for(var t=arguments.length,e=new Array(t),n=0;n<t;++n)e[n]=arguments[n];return e}var u=t("ee"),d=t(33),p="nr@original",l=Object.prototype.hasOwnProperty,h=!1;e.exports=r,e.exports.wrapFunction=s,e.exports.wrapInPlace=c,e.exports.argsToArray=f},{}]},{},["loader",2,17,5,3,4]);</script><style type="text/css">
.jsEnabled {display:none;}
</style>
<META NAME= "robots" CONTENT = "NOARCHIVE, NOFOLLOW, NOINDEX">
<META NAME="description" CONTENT="TEST 02 - Elsevier's Scopus, the largest abstract and citation database of peer-reviewed literature. Search and access research from the science, technology, medicine, social sciences and arts and humanities fields.">	
<link rel="SHORTCUT ICON" href="/static/proteus-images/favicon.ico?ver=1.0">
<script src="https://components.scopus.com/www/components/jquery.min.js"></script>
<script src="https://components.scopus.com/www/components/jquery-ui.min.js"></script>
<script type="text/javascript" src="/gzip_N1660835974/bundles/masterjquery.js" ></script>
<script language="javascript" type="text/javascript">
window.$ = window.jQuery;
</script>
<script>
var feMetricUrl = new Object();
feMetricUrl = 'https://rum.scopus.com';
</script>
<link rel="stylesheet" type="text/css" media="all" href="/gzip_1596155880/bundles/ScopusMasterLayout.css" />
<link rel="stylesheet" type="text/css" media="all" href="/gzip_N1361801193/bundles/d3charts.css" />
<script type="text/javascript" src="/gzip_N888159981/bundles/RecordPageTopMaster.js" ></script>
<script type="text/javascript">
var MendeleybookmarkletUrl = '';
var isPreviewPage = true;
var selectedNav = "";
var isRegisteredPreview = 'true' === 'false';
</script>
<script type="text/javascript">
/* Deprecated - use isIndividuallyAuthenticated instead */ isLoggedInUser = 'true' === 'false';
isShibUser = ('GUEST'==="SHIBBOLETHANON")?true:false;
var isIndividual = 'true' === 'false';
isIndividuallyAuthenticated = isLoggedInUser || isIndividual;
</script>
<script type="text/javascript">
var isIDPlusEnabled = 'true' === 'true';
var isCarsForIDPlus = 'true' === '';
var isCARSBulkRegisterPage = 'true' === '';
var isCARSRegisterPage = 'true' === '';
var isCARSPathChoicePage = 'true' === '';
var isCARSLoginFullPage = 'true' === '';
</script>
<script>
var smapiVars = {};
smapiVars["FEATURE_SET"] = "FEATURE_DOCUMENT_RESULT_BETA_RELEASE:0,FEATURE_DOCUMENT_RESULT_MICRO_UI:0,FEATURE_USE_MENDELEY_DOWNLOAD:1,FEATURE_SQT_LIBRARY_ENABLE:1,FEATURE_AFW_REQUEST_MICRO_UI:1,FEATURE_SERVER_SIDE_AB_TEST:0,FEATURE_AFFILIATION_DATE_RANGE:0,FEATURE_IDENTITY_PLUS:1,FEATURE_SEARCH_SYNTX_MSG:1,FEATURE_BREAK_UP_MONOLITH:1,FEATURE_REFWORKS_ENCRYPTION:1,FEATURE_DEMARCATED_LAYER:1,FEATURE_BUM_HEADER:1,FEATURE_IE11_SUNSET:1,FEATURE_USER_DASHBOARD_MICRO_UI:1,FEATURE_UN_SDG:1,FEATURE_PEOPLE_FINDER:0,FEATURE_HOMEPAGE_MICRO_UI:1,FEATURE_VIEWS_COUNT:1,FEATURE_NEW_SCIVAL_TOPICS:1,FEATURE_NEW_REAXYS_SECTION:1,FEATURE_NEW_SOURCE_INFO:1,FEATURE_NEW_REFERENCES_SECTION:0,FEATURE_NEW_METRICS_SECTION:1,FEATURE_NEW_CITED_BY_SECTION:0,FEATURE_NEW_RELATED_ARTICLE_DATA:0,FEATURE_NEW_RELATED_DOCUMENTS_SECTION:0,FEATURE_NEW_RECORD_PAGE_PAGINATION_AND_BACKTOTOP:0,FEATURE_NEW_QUICK_LINKS:0,FEATURE_NEW_OUTWARD_LINKS:0,FEATURE_NEW_ABSTRACT_CORRESPONDENCE:0,FEATURE_NEW_RAW_DOCUMENT_VIEW:0,FEATURE_NEW_MAIN_SECTION:1,FEATURE_DOC_DETAILS_TOOLBAR:1,FEATURE_VIEW_AT_REPOSITORY:1,FEATURE_VIEW_PDF:0,FEATURE_FULL_TEXT_OPTIONS:1,FEATURE_STENCIL:1,FEATURE_SEARCH_HISTORY_MICRO_UI:1,FEATURE_HOMEPAGE_SAVED_SEARCHES:1,FEATURE_HOMEPAGE_SEARCH_HISTORY_SAVE_SEARCH_ACTION:1,FEATURE_AFW_MVP:1,FEATURE_NEW_DOC_DETAILS_EXPORT:0,FEATURE_NEW_DOC_DETAILS_PRINT:0,FEATURE_NEW_DOC_DETAILS_EMAIL:0,FEATURE_NEW_DOC_DETAILS_SAVE_PDF:0,FEATURE_NEW_DOC_DETAILS_SAVE_TO_LIST:0,FEATURE_NEW_DOC_DETAILS_ORDER_DOCUMENT:1,FEATURE_DOC_DETAILS_MORE_BUTTON:0,FEATURE_NEW_DOC_DETAILS_CREATE_BIBLIOGRAPHY:0,FEATURE_NEW_OPEN_ACCESS_TYPES:1,FEATURE_HOMEPAGE_SCOPUS_USE_CASES:1,FEATURE_HOMEPAGE_CONTENT_METRICS:1";
smapiVars["SERVICE_DOMAIN"] = "https://www.scopus.com/api";
smapiVars["fe.id"] = "www";
smapiVars["RESULT_SET_SIZE"] = "2000";
smapiVars["AFW_AUTHOR_MAX_SELECT_COUNT"] = "8";
smapiVars["MAX_SOURCE_RESULTS_DISPLAYED"] = "1000";
smapiVars["PLATFORM_BRIDGE_REQUEST_PROXY_URL"] = "https://www.scopus.com/serverProxy/request.uri";
smapiVars["APP_DOMAIN"] = "www.scopus.com";
smapiVars["AFW_PARTNER_ID"] = "fpDjhFsn";
smapiVars["AFW_MICRO_UI_BUNDLE_BASE_URL"] = "https://components.scopus.com/www";
smapiVars["MICRO_UI_BUNDLE_BASE_URL"] = "https://components.scopus.com/www";
smapiVars["STYLEGUIDE_BUNDLE_BASE_URL"] = "https://components.scopus.com/www";
smapiVars["SCIVAL_AUTHOR_EXPORT_URL"] = "/import/externalResearcher";
smapiVars["SCIVAL_SERVER_DOMAIN"] = "www.scival.com";
smapiVars["USE_DOCUMENT_PREPRINTS"] = "true";
smapiVars["MENDELEY_DOMAIN"] = "https://www.mendeley.com";
smapiVars["SV_TOPIC_WEB_LINK"] = "https://scival.com/viewtopic?id=";
smapiVars["AUTHORS_WITH_AUTHOR_ID_DISABLED"] = "24476171900";
smapiVars["ORCID_DOMAIN"] = "orcid.org";
smapiVars["USE_AWARDED_GRANTS"] = "false";
smapiVars["EXPERT_LOOKUP_DOMAIN"] = "https://api.expertlookup.com";
smapiVars["EXPERT_LOOKUP_WORKSPACE_ID"] = "3e04d3c0-083a-eb11-8172-db2c1b088b85";
smapiVars["AWARDED_GRANTS_API_KEY"] = "98FBCD4C57A1F64F9BB9A3B38B993";
smapiVars["DOC_DETAILS_SERVICE_DOMAIN"] = "https://documents-facade.scopus|.|com";
smapiVars["LISTS_API_SERVICE_DOMAIN"] = "";
smapiVars["API_GATEWAY_DOMAIN"] = "https://api.scopus.com";
smapiVars["VIEW_PDF_SRC"] = "https://static.mendeley.com/view-pdf-component/0.5.0/dist/view-pdf-element.js";
</script>
<title>
Scopus preview - 
Scopus - Document details - An efficient highly parallelized ReRAM-based architecture for motion estimation of HEVC
</title>
<link rel="dns-prefetch" href="https://assets.adobedtm.com"/>
<link rel="preconnect" href="https://assets.adobedtm.com" crossorigin />
<link rel="dns-prefetch" href="https://smetrics.elsevier.com" />
<link rel="preconnect" href="https://smetrics.elsevier.com" crossorigin />
<link rel="dns-prefetch" href="https://mboxedge37.tt.omtrdc.net" />
<link rel="preconnect" href="https://mboxedge37.tt.omtrdc.net" crossorigin />
<link rel="dns-prefetch" href="https://dpm.demdex.net" crossorigin />
<link rel="preconnect" href="https://dpm.demdex.net" crossorigin />
<link rel="dns-prefetch" href="https://elsevierlimited.tt.omtrdc.net" />
<link rel="dns-prefetch" href="https://elsevier.demdex.net" />
<link rel="dns-prefetch" href="https://content.pendo.scopus.com" />
<link rel="preconnect" href="https://content.pendo.scopus.com" crossorigin />
<link rel="dns-prefetch" href="https://app.pendo.io" />
<link rel="preconnect" href="https://app.pendo.io" crossorigin />
<link rel="dns-prefetch" href="https://pendo-io-static.storage.googleapis.com" />
<link rel="dns-prefetch" href="https://js-agent.newrelic.com" />
<link rel="preconnect" href="https://js-agent.newrelic.com" crossorigin />
<link rel="dns-prefetch" href="https://bam.nr-data.net" />
<link rel="dns-prefetch" href="https://rum.scopus.com" />
<link rel="dns-prefetch" href="https://secure.adnxs.com" />
<link rel="dns-prefetch" href="https://components.scopus.com" />
<link rel="preconnect" href="https://components.scopus.com" crossorigin />
<link rel="dns-prefetch" href="https://cdn.elsevier.io" />
<link rel="preconnect" href="https://cdn.elsevier.io" crossorigin />
<script type="text/javascript">
var pageData = {
page: {},
visitor: {}
};
document.addEventListener('pageDataUtilReady', function () {
pageDataUtil.loadPageDataObject("false", "2.44.73.40", "en_US", "", "13341253", "ae:ANON::GUEST:",
"", "278641", "", "Scopus Preview", "700DB85889C256DC6548C555F9984E93.i-061e2a7b9e7281261");
});
</script>
<script type="text/javascript" src="/gzip_2040108700/bundles/SiteCatalystTop.js" ></script>
<script defer="true" src="//assets.adobedtm.com/4a848ae9611a/8abdb8d26bcc/launch-5ed435cee62b.min.js" ></script>
<script type="text/javascript">
var isBreakUpMonolith = 'true' === 'true';
var isBreakUpMonolith = 'true';
</script>
<script type="text/javascript">
var ScopusUser = {
webUserId: 13341253,
idToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImhHcWNQcDdZa0ZTYWp1eVYxLXpaVVNRMGpXUSJ9.eyJzdWIiOiIxMzM0MTI1MyIsImF1ZCI6IlNDT1BVUyIsImp0aSI6ImJ4RExxeFI1dVJERUxnTU5GS0dhQTciLCJpc3MiOiJodHRwczovL2lkLmVsc2V2aWVyLmNvbSIsImlhdCI6MTYzODExOTIzMiwiZXhwIjoxNjM4MTE5NTMyLCJwYXRoX2Nob2ljZSI6ZmFsc2UsImFuYWx5dGljc19pbmZvIjp7ImFjY2Vzc1R5cGUiOiJhZTpBTk9OOjpHVUVTVDoiLCJhY2NvdW50SWQiOiJhZToyNzg2NDEiLCJhY2NvdW50TmFtZSI6ImFlOlNjb3B1cyBQcmV2aWV3IiwidXNlcklkIjoiYWU6MTMzNDEyNTMifSwiaW5zdF9hY2N0X25hbWUiOiJTY29wdXMgUHJldmlldyIsInVzZXJfaW5mb19leHAiOiIxNjM4MTIwMTMyIiwiaW5zdF9hY2N0X2ltYWdlIjoiaHR0cDovL2lkLmVsc2V2aWVyLmNvbS9hc3NldHMvaW1hZ2VzL2Vsc2V2aWVyL2RlZmF1bHRfaW5zdGl0dXRpb24uSlBHIiwiaW5kdl9pZGVudGl0eSI6IkFOT04iLCJpbnN0X2Fzc29jIjoiR1VFU1QiLCJpbnN0X2FjY3RfaWQiOiIyNzg2NDEiLCJhdXRoX3Rva2VuIjoiNmI2YWU1ZDk2ODNiMDQ0YmUzMDkzYTczODUyN2JlZGU4ZjJiZ3hycWIifQ.LTB22KYz5GJ2FHMQcJJYcnLaaIPhasFT1hM4fGv3McU-cyUicOvH10u2sSlTr4OnGH0riIz-1a6xzKbA5g0sk26WYBnFrtyXhZk6YOGMSW7lg9rUPS-m4C7eoy6S6p2qPucneIRK0kFZKKu0SbWK10E4BWHb4Zns7svm52LmPBD0FLboHVH5s0oarqiBjNGtI_1SEtL1OnHc0gw0fPbheQrrfO7eAh8yDeDWpmUZqdigxiYlQWXzMhO-k9duYoi15XEfosoUw4Io5zCb_45gFwzDrL1pycyk2ah6iCoBrJ7BC3P-lHPqJenFxoZxOP7maqOBH9sGwhGyLqxDh8CIpw",
firstName: "",
lastName: "",
email: "",
accessTypeAA: "ae:ANON::GUEST:",
department: {
departmentId: "289839",
departmentName: "Scopus Preview",
},
account: {
accountId: "278641",
accountName: "Scopus Preview"
},
isIndividual: "true" === "false",
isSubscribed: "true" === "false",
pathChoiceExists: "true" === "false",
fenceIds: [40, 51, 53, 54, 55, 14371, 14373, 14374, 14376, 18365, 24379, 24380, 24382, 29365, 31365, 31366, 34371, 34377, 34380, 41382, 41383, 41384, 41385, 41386, 41387, 41388, 44374, 55371, 60386, 60395, 95379, 103379, 104379, 106386, 107379, 108379, 110383, 111379, 116380, 120380, 120381, 434380, 464380],
fenceNames: "[ENABLE_INT_POINT_BELOW_ABS_PRIVILEGED, ENABLE_SCOPUS_REMOTE_ACCESS_POPUP, BLOCK_FACEBOOK_LOGIN, ENABLE_ALIAS_RESOLVE_TOMBSTONES, ALLOW_PEOPLE_FINDER, BLOCK_ALL_SOCIAL_LOGINS, BLOCK_LINKED_LOGIN, ENABLE_TERM_ANALYZER, PROMOTE_SECONDARY_RESULTS, BLOCK_OPENID_LOGIN, ALLOW_PDF_DOWNLOAD, ENABLE_CHINESE_SIMPLIFIED, BLOCK_YAHOO!_LOGIN, ENABLE_AFFILIATION_HIERARCHY, isEnableXtrnlSubscribedEntitlements, ENABLE_SITE_CATALYST_TRACKING, RECORD_PAGE_TOOLBAR_ENABLE, FENCE_OPENATHENS_LOGIN_LINK, ENABLE_RUSSIAN, DDM_DISABLE_QUOSA, ENABLE_REAXYS_SUBSTANCES, DISABLE_SCOPUS_REMOTE_ACCESS_POPUP, ENABLE_AFFIL_HIERARCHY_IND_COUNT_DISPLAY, FENCE_INSTITUTIONAL_SHIB_LOGIN_LINK, BLOCK_TWITTER_LOGIN, ALLOW_SCIVAL_NAVIGATION, ENABLE_INTEGRATION_POINT_BELOW_ABSTRACT, DDM_POPUP_FLAG_FIREFOX, DDM_ENABLE_INLINE_INSTALLATION_CHROME, DDM_ENABLE_MACHINE_LEARNING_ALGORITHM, ENABLE_PENDO, ALLOW_MORE_DOCUMENTS_TAB, ALLOW_MANRA, DDM_POPUP_FLAG_CHROME, ENABLE_CHINESE_TRADITIONAL, AUTHDETAILS_PAGE_TOOLBAR_ENABLE, SEARCH_TERM_HIGHLIGHTING, ALLOW_AUTHOR_AFFILIATION_PROFILING, ALLOW_MORE_AFFILIATIONS, BLOCK_UN_TO_EMAIL_MIGRATION_AT_LOGIN, ALLOW_AFFILIATION_PROFILING, ENABLE_SCIVAL_TOGGLE, isAdminUser, ALLOW_SHIBBOLETH_AUTHENTICATION, ENABLE_JAPANESE_INTERFACE_FOR_SCOPUS, BLOCK_GOOGLE_LOGIN]".replace(/(\[)|(\])|( )/g, "").split(","),
adminRoles: "[]".replace(/(\[)|(\])|( )/g, "").split(","),
userPreferences: {"ENABLE_REFWORKS_ICON":"false"},
userLocale: "en_US".replace("_","-")
};
</script>
<script type="text/javascript">
var PlatformData = {
sessionId: "700DB85889C256DC6548C555F9984E93.i-061e2a7b9e7281261",
pageId: "recordpage",
sessionOverrides: {}
};
PlatformData.sessionOverrides = PlatformData.sessionOverrides || {};
</script>
<script
src="https://components.scopus.com/www/experimental-components/experimental/experimental.js"
></script>
<script src="https://components.scopus.com/www/components/react.production.min.js"></script>
<script src="https://components.scopus.com/www/components/react-dom.production.min.js"></script>
<script src="https://components.scopus.com/www/pbj-scopus-monolith/pbj-scopus-monolith.js"></script>
<script src="https://components.scopus.com/www/components/custom-elements-es5-adapter.js"></script>
<script src="https://components.scopus.com/www/components/webcomponents-loader.js"></script>
<script src="https://components.scopus.com/www/components/scopus-components.js"></script>
<link rel="stylesheet" type="text/css" media="all" href="https://components.scopus.com/www/stylesheet/scopus-stylesheet.css">
<script src="https://components.scopus.com/www/micro-ui-injector/micro-ui-injector.js"></script>
</head>
<body onLoad="onBodyLoad();" data-page-type="" id="resultsPagePreview">
<div class="container-fluid">
<div class="row-fluid mainRow" id="rowMain">
<div id="contentContainer" class="col-md-12 fullHeightHack content">
<header id="gh-cnt" class="highlight-none">
<a class="anchor gh-skip-navigation" href="#skip1"><span class="anchor-text">Skip to main content</span></a>
<a name="top" id="top"></a>
<input type="hidden" id="txGid" value="700DB85889C256DC6548C555F9984E93.i-061e2a7b9e7281261:5"/>
<script>
var prs = {rt: function(label,ts) { this[label]=(ts||ts==0?ts:new Date().getTime());}};
var continuePcrRecording = true;
var PCR_SCR_ATTRIBUTE_NAME = 'data-reportRenderTime';
function recordRendering (primaryContentElements, mutations) {
if (continuePcrRecording) {
primaryContentElements.forEach(function (element) {
for (var i = 0; i < mutations.length; i++) {
var changedNode = mutations[i].target;
if (element === changedNode || element.contains(changedNode)) {
var contentSignificance = element.getAttribute(PCR_SCR_ATTRIBUTE_NAME);
prs.rt(contentSignificance);
}
}
});
}
};
MutationObserver = window.MutationObserver || window.WebKitMutationObserver;
var domChangeListener = new MutationObserver (function (mutations) {
// fired when the document changes - look for PCR / SCR changes
recordRendering (document.querySelectorAll('[' + PCR_SCR_ATTRIBUTE_NAME + '="pcr"]'), mutations);
recordRendering (document.querySelectorAll('[' + PCR_SCR_ATTRIBUTE_NAME + '^="scr"]'), mutations);
});
domChangeListener.observe (document, {childList: true, attributes: true, characterData: true, subtree: true});
</script>     
<div class="gh-lib-banner gh-lb-dominant hidden">
<input type = "hidden" id="bannerHeaderCheck" value="notAvailable">
</div>
<div id="gh-main-cnt" class="u-flex-center-ver u-position-relative gh-sides-padding">
<a id="gh-branding" class="u-flex-center-ver" href="/home.uri?zone=header&amp;origin=recordpage" aria-label="Scopus (opens in new window)"><span class="anchor-text"><svg xmlns="http://www.w3.org/2000/svg" role="img" version="1.1" width="54" height="48" viewBox="0 0 54 48" enable-background="new 0 0 54 48" xml:space="preserve" class="gh-logo" focusable="false" aria-hidden="true"><title id="gh-elsevier-logo">Elsevier logo</title><path xmlns="http://www.w3.org/2000/svg" d="M15.04 22.48c.58-.3.91-.68.97-1.11.09-.4-.26-.36-.35-.28-.32.26-.74.66-1.14.93a4.19 4.19 0 0 0-.56.07c-.05-.02-.1-.1-.08-.12.37-.23.77-.46 1.03-.77.17-.2.1-.37-.23-.41-.37-.06-.8.1-1.13.39-.57.48-.84 1.55-.9 2.23-.11.12-.15.12-.27.17a2.83 2.83 0 0 1 .3-2.2c.05-.14.01-.37-.2-.43-.11-.05-.17-.02-.22.1-.28.81-.77 1.15-1.45 1.66-.06.05-.68.39-.83.48-.54-.17-1.15-1.17-1.66-1.93a.47.47 0 0 0-.45-.2c-1.11.13-1.94.05-2.54-.83-.27-.36-.39-.43-.72-.11-.38.39-.75.69-1.12 1-.2.13-.34.09-.42-.16a7.04 7.04 0 0 1-.37-2.35c0-.33.04-.84.05-1.16a2.87 2.87 0 0 0-.71-.48C1.1 16.53 0 15.5 0 15.17c0-.07.06-.25.32-.4 1.36-.6 2.43-.1 2.83-.16.35-.03.46-.28.4-.77a2.7 2.7 0 0 1 .88-2.42c.46-.5.43-1.04.32-1.65-.06-.28-.12-.35-.4-.32-.75.16-.71.18-.85.83-.08.42-.23.85-.47 1.2-.3.47-.7.87-1.05 1.29-.12.17-.26.17-.44.05A2.44 2.44 0 0 1 .6 9.74c.12-.32.3-.67.49-.84l.1.03c0 .06 0 .13-.03.19l-.22.69c-.28.8-.23 1.57.3 2.26.18.22.34.29.44.19.34-.35.7-.67 1-1.05.36-.42.47-.93.36-1.46a9.27 9.27 0 0 0-.1-.6c-.05-.3-.18-.27-.32-.1-.33.34-.49.61-.65 1.23a.48.48 0 0 1-.15.23.65.65 0 0 1-.08-.33c.1-.56.22-1.05.58-1.41.23-.26.17-.4-.05-.46-.8-.22-1.49-.45-2.1-1.2a.67.67 0 0 1-.16-.4c0-.16.08-.28.26-.41.7-.48 1.42-.99 2.11-1.14l.57-.09C2.62 4.1 2.3 2.9 2.1 1.81c0-.23.16-.84.25-1.02.15-.02.34-.02.62.08.32.12.62.32.92.5.29.17.37.17.56-.09.32-.47.7-.89 1.19-1.14.17-.08.28-.14.39-.14.16 0 .28.1.44.34.27.43.5.93.74 1.4.28.47.58.5 1 .19.48-.32.91-.63 1.35-.99.4-.32.52-.32.97 0 .4.27.83.78 1.05 1.23.04.08.97-.42 1-.64.04-.4.15-.7.3-1.02.2-.4.47-.51.73-.51.23 0 .67.26 1 .47l.7.4c.37.2.39.2.61-.13.25-.4.47-.47.99-.47 1.02 0 1.95.26 2.66 1.1.13.17.25.2.41.04.29-.3.56-.41 1.1-.41h.95c.21 0 .35.01.55-.26.26-.38.55-.74.97-.74.67 0 1.55.94 2 1.75l-.24.12a4.76 4.76 0 0 0-1.75-1.1c-.4-.17-.58 0-.53.42.06.5.19 1 .3 1.49.15.46.55.86 1 1.05l.02.17c-.07-.02-.16-.05-.28-.06a2.04 2.04 0 0 0-.87.03c-.56.14-1.07.38-1.57.6-.25.12-.25.18-.04.39.28.26.54.56.86.79.9.68 1.8.54 2.79.16l.46-.16c.42-.08.87.1 1.21.4l1 .91c.32.3.4.58.2.96-.33.64-.7 1.3-1.06 1.92a.62.62 0 0 1-.23.24c-.12-.1-.23-.33-.11-.53.44-.56.72-1.34.82-1.87.09-.52-.19-.77-.68-.64l-.42.16c-.18.07-.24 0-.28-.18-.07-.16-.15-.26-.2-.43-.12-.28-.28-.37-.52-.19-.42.26-.71.52-.96.93a3.53 3.53 0 0 0-.4 2.17c.01.1.09.16.14.22a.4.4 0 0 0 .19-.2c.13-.32.21-.64.27-1 .03-.2.06-.54.12-.72.05-.1.2-.06.28-.03.21.08.3.19.33.34-.42.53-.68 1.32-.96 2.02-.05.1-.23.22-.34.22-.13-.01-.3-.12-.34-.23-.17-.45-.47-.89-.59-1.34-.3-1.51-.4-1.6-.46-1.69a9.32 9.32 0 0 0-2.26-2.57c-.07-.07-.1-.16-.16-.22.08-.04.15-.11.23-.14l.88-.34c.38-.13.64-.37.8-.71.1-.2.25-.4.36-.59.16-.28.14-.43-.07-.67a1.44 1.44 0 0 0-1.47-.45c-.57.14-.92.54-1.1 1.08-.07.2-.09.4-.09.61 0 .3.1.37.35.21.41-.34.88-.62 1.33-.92-.2.48-.5.85-.92 1.12-.28.18-.55.32-.83.48-.24.17-.32.1-.4-.18-.06-.3-.03-.6-.03-.9A1.97 1.97 0 0 0 18 .9l-.74-.3a.7.7 0 0 0-.28.64c.06.8.18 1.4.32 2.08.13.28.44.4.75.45.19-.1.34-.36.09-.76l-.37-.72c-.03-.1-.08-.18-.1-.28-.02-.04-.04-.1 0-.15.04.02.09.02.18.11.46.5.84 1.07 1.08 1.78.1.28 0 .39-.26.42-.62.05-1.24.1-1.84.22-.97.2-1.3.56-1.65 1.21-.15.28-.23.46-.24.62.35.03.74.01 1.3-.09a3.7 3.7 0 0 0 1.34-.44c.26-.3.46-.61.85-.89.12-.08.38-.11.48-.05a5.12 5.12 0 0 1 2.56 2.5c.04.12.15.23.2.32.11.13.07.18-.09.17-.22-.06-.47-.12-.67-.2-.08-.04-.15-.17-.19-.26l-.2-.58c-.11-.23-.17-.23-.4-.15l-.45.19c-.17.06-.25.09-.29-.1-.05-.28-.13-.55-.2-.8-.1-.31-.22-.38-.52-.23a2.8 2.8 0 0 0-1.25 1.28c-.26.54-.36.9-.12 2.34.06.28.28.62.54.7l2.42.75c.12 0 .2.14.05.24-.3.2-.63.38-.94.57-.15.1-.24.05-.31-.1a1.9 1.9 0 0 0-1-.9c-.16-.08-.24-.05-.3.12-.12.26-.24.64-.34.91-.06.14-.12.2-.28.1-.69-.58-1.41-.61-2.26-.5-.42.06-.61.17-.51.59.23 1.16 1.25 1.7 2.11 2.07a.59.59 0 0 0 .3-.23c.13-.19.1-.37-.1-.54-.33-.29-1.04-.62-1.39-.94-.06-.19.34-.35.56-.21.63.28 1.11.8 1.62 1.3.26.25.38.22.46-.14.1-.28.19-.7.27-1.03.05-.16.15-.1.19 0 .06.24.14.55.14.83 0 .3-.1.56-.19.84-.05.28-.27.43-.5.37-1.07-.1-1.88.4-2.54.42-.15 0-.47-.1-.62-.17-.24-.18.1-.2.18-.2.28 0 .7 0 .92-.12-.06-.1-.18-.2-.29-.27-.6-.49-1.26.03-1.77.29-.3.16-.31.2-.19.5.19.46.52.64.98.64.26 0 .51.04.78.03a.93.93 0 0 0 .37-.1c.35-.12.65-.3.97-.43.4-.06.38.14.2.33l-.79.49c-.08.04-.17.07-.23.13-.11.08-.25.22-.23.28.05.1.17.21.3.26.37.11.74.17 1.1.25.57.11.81-.05.97-.62.08-.31.19-.6.34-.9.09-.18.2-.2.31-.02.4.73.8 1.44 1.22 2.14.08.16 0 .24-.15.22-.9-.1-1.8-.2-2.7-.34-.3-.04-.6-.24-.9-.37a3.2 3.2 0 0 0-1.47-.39c-.45 0-.72.28-.56.72.2.62.44 1.26.7 1.89.17.42.59.58 1.04.62.22.03.14-.12.09-.22l-.75-1.61c-.1-.27-.02-.31.13-.09.41.62.84 1.2 1.21 1.83.23.39-.05.5-.33.5-.94-.09-1.8-.41-2.7-.76-.1-.03-.22-.23-.2-.31.17-.16.38 0 .56 0 .22.04.3-.06.22-.22a.86.86 0 0 0-.45-.42 5.56 5.56 0 0 0-1.34-.49c-.28-.05-.47.07-.41.32.07.32.22.61.39.95.05.1.05.22.02.31-.17 0-.37.05-1.07-.25-.7-.29-.7-.48-.7-1.34 0-.18 0-.55.03-.71.07-.24.1-.28.3-.25.73.15 1.47.35 2.18.7.22.09.42.27.65.36.18.06.27 0 .25-.19a.65.65 0 0 0-.41-.56c-.77-.32-1.53-.69-2.3-1-.18-.07-.24-.15-.14-.35.3-.5.46-1.09.33-1.72 0-.05.06-.11.12-.17a.47.47 0 0 1 .34.43c.06.54.42.91.77 1.28.72.74.72.74 1.72.5.13-.02.23-.11.34-.19-.1-.07-.19-.2-.32-.26-.18-.1-.39-.15-.6-.23-.66-.25-.86-.73-.64-1.43a24.2 24.2 0 0 0 .63-1.95c.07-.3.1-.63.11-.96 0-.18.14-.18.25-.2l1.46-.1c.37-.02.8-.46.94-.71l.49-.8a1.01 1.01 0 0 0-1-.14c.11-.36.26-.7.26-1.05 0-.23-.05-.5-.15-.73-.1-.22-.25-.26-.44-.12-.35.3-.65.65-.76 1.14-.11.48-.31 1.36-.21 1.52.22-.17.52-.39.71-.45.1-.02.26 0 .34.06-.33.3-.72.56-1.08.75-.27.12-.37-.04-.33-.3.32-1.21.12-1.3.02-1.9a6.98 6.98 0 0 0-.9-2.09c-.13-.16-.34-.1-.61.23-.45.54-.56 1-.8 1.92l-.05.19-.32.28-.14-.16c.12-.43.25-.68.25-1.02 0-.42-.18-.72-.46-.95-.35.22-.66.5-.75.65a.92.92 0 0 0 .02 1.06c.36.62.3 1.25.19 1.98l-.1-.03a8.26 8.26 0 0 0-.25-1.41c-.08-.3-.27-.55-.43-.8-.16-.24-.3-.74-.17-.95.19-.32.77-.5 1-.79.23-.35.12-.5.41-.82.18-.21.56-.65.84-.5.84.42 1.73.51 2.67.39.4-.06.49-.05.6-.45.1-.33.14-.1.47-.22.57-.19.87-.7.87-1.3 0-.34.02-.68-.02-1.03-.03-.3-.14-.58-.23-.85-.08-.18-.2-.14-.28-.01-.23.3-.63.9-.84 1.24l-.33.6c-.05.16-.13.2-.27.05a3.03 3.03 0 0 0-.49-.43c-.3-.2-.33-.81-.21-1.2.13.07.22.22.3.42.12.26.17.56.38.65.27.14.54-.42.5-.8a.6.6 0 0 0-.1-.28C14.77.83 14.1.4 13.68.4c-.48 0-.8.6-.6 1.16.06.15.12.26.2.4-1.15.6-1.46.99-1.92 1.52-.05.07-.1.25-.05.34.21.41.13.55-.31.8-.2.1-.46.21-.67.27-.03.02-.08-.02-.14-.03a.28.28 0 0 1 .06-.11c.28-.63.58-1.22.84-1.86.28-.68-.33-1.3-.88-1.52-.36-.14-.65-.06-.9.27-.6.72-.73 1.9.12 3.02.13.19.1.26-.12.28-.49.04-.97.1-1.47.12a1.95 1.95 0 0 0-1.5.8c-.1.16-.25.2-.41.09-.38-.24-.7-.18-1.1.05a4.9 4.9 0 0 1-1.84.62c-.07 0-.15-.06-.2-.12.03-.05.09-.13.23-.18.4-.1.72-.2 1.08-.3.06-.06.09-.22.13-.29-.07-.1-.2-.21-.31-.22a4.91 4.91 0 0 0-1.02 0c-.88.08-1.53.5-2.09 1.01-.26.26-.28.53 0 .75.65.46 1.35.93 2.23.7.62-.2 1-.66 1.4-1.1.33-.36.74-.77 1.13-.33.57.65.68 1.33 1.61 1.21.36-.04.5-.16.6-.5l.07-.4c.03-.25-.1-.38-.36-.28-.17.06-.31.15-.48.19-.1.02-.2-.01-.3-.04 0-.1.02-.22.08-.28.17-.11.38-.2.56-.29.7-.18 1.36-.32 2.05-.46a.6.6 0 0 1 .21 0l-.06.2c-.25.57-.59.75-1.04.86-.06.05-.17.18-.25.44a2.6 2.6 0 0 0 .5 2.57c.22.25.5.46.79.63.48.28.61.75.79 1.24.11.31.23.62.37.92-.39.28-.71.6-.94 1-.1.1-.17.06-.22-.05l-.27-.66c-.24-.58-.65-.96-1.3-1.01-.24-.02-.47-.06-.7-.09-.26-.02-.44.1-.53.4-.37 1.12-.02 2.09 1.05 2.68-.86.22-1.56.56-2 1.16.04.17.3.39.64.53.84.34 1.7.3 2.49.04.46-.17.48-.06.56.28l.24 1.04c.05.25-.09.39-.32.3-1.28-.37-2.8-.98-3.9-1.72a5.77 5.77 0 0 0-1.43-.43c-.17.1-.31.22-.46.2a8.16 8.16 0 0 1-1.38-.5c-.22-.12-.23-.19.05-.23.43 0 1.21.19 1.61.19.4-.12.4-.19.15-.35-.94-.75-2-.57-3.05-.4-.38.08-.4.26-.19.56 1.36 1.53 2.36 1.52 4.02 1.33 1.08.68 2.16 1.3 3.42 1.7l.02.18c-.9.06-1.75.21-2.56.38-.41.08-.5.33-.27.7.33.54.86.8 1.47.94.43.13.87.1 1.3-.08.5-.23.83-.65 1.3-.91.28-.15.75.02 1.03.1.42.13.83.32 1.27.38 1.73.24 3.57.4 5.18.26-.2.38-.33.73-.46 1.1-.09.22-.01.27.22.25.34 0 .67-.2.81-.46.03-.26.03-.5.08-.76.03-.14.07-.33.16-.35.07-.01.47.06.42.17-.14.4-.26.71-.43 1.11.15.03.31.06.47.06.65 0 .75-.1.72-.74a2.7 2.7 0 0 0-.08-.55c-.08-.49.02-.82.44-1.05.4-.2.84-.14 1.25-.09a7.63 7.63 0 0 1 3.52 1.6c.51.4.07.97.07 1.74 0 .49-.1.99-.09 1.24.02.4.19.61.34.73.45.34.94.68 1.41 1 .49.3.73.5 1.3.7.14.04.3.04.42.03a1.49 1.49 0 0 0-.03-.6c-.28-.19-.67-.28-1.12-.43-.08-.03-.17-.14-.21-.24a.84.84 0 0 1 .32-.19l.73-.15c.05-.02.15-.1.18-.16-.03-.07-.09-.15-.14-.16-.3-.1-.7-.16-.94-.22-.1-.03-.14-.12-.22-.2a.91.91 0 0 1 .36-.19l.74-.09a.4.4 0 0 0 .2-.12c-.03-.07-.07-.17-.13-.2-.15-.07-.92-.35-1-.43-.07-.07-.08-.16-.11-.28.1-.05.25-.1.35-.08l.62.01a.4.4 0 0 0 .26-.12.56.56 0 0 0-.17-.2c-.3-.18-.41-.3-.75-.46-.15-.1-.23-.18-.26-.38.12-.03.28-.05.42-.02.23.06.35.13.59.17.06.02.18-.01.26-.03a.26.26 0 0 0-.07-.18c-.09-.1-.21-.14-.32-.23-.17-.11-.46-.39-.63-.53-.05-.04-.02-.18-.03-.23h.18c.23.05.56.15.75.18.33.06.29.01.15-.23-.2-.18-.42-.3-.62-.52-.05-.07-.03-.2-.03-.3.32.02.65.17.94.17.02-.05.04-.1.04-.17a4.75 4.75 0 0 0-.23-.55c-.08-.22.03-.3.19-.24.22.1.57.28.81.34.1.03.22 0 .34-.05-.11-.22-.24-.4-.4-.62-.02-.06 0-.14.06-.2.38.1.77.37 1.11.46.16.06.26 0 .3-.08.03-.09-.25-.37-.23-.48 0-.07.02-.1.06-.14l.65.2c.48.13.75.17.82.06.1-.03 0-.25-.15-.51.38-.03.42.15 1.08.28.25.06 3.74-.43 5.16-.87.56-.18 1.2-.46 1.73-.7l.17.03c.4 1.03 1.24 1.34 2.1 1.44l.8.05c.25 0 .3-.06.23-.31a4.47 4.47 0 0 0-1.3-2.2l-.55-.35c-.24-.25-.21-.54.05-.74.44-.3 1.04-.59 1.49-.87.65-.41.87-1.06.83-1.74a.58.58 0 0 0-.24-.45.73.73 0 0 0-.33.23c-.18.37-.36.68-.51 1-.13.07-.26.1-.37.15-.03-.12-.1-.23-.06-.34.2-.58.7-1.04 1.04-1.41.14-.24-.07-.33-.17-.31-.37.08-.78.22-1.06.54a7.5 7.5 0 0 0-1.47 3.04c-.06.29-.19.4-.28.66-.12.41-.11.5.24.76.52.33 1.19.64 1.67.98.2.22.1.3-.11.3a1.93 1.93 0 0 1-1.4-.63c-.18-.2-.34-.34-.56-.37-.5.14-1.2.34-1.72.41-.3.06-.53.18-.69.46-.1.23-.31.28-.56.19-.37-.1-.74-.24-1.11-.32a4.19 4.19 0 0 0-1.86.13c-.25.04-.31-.02-.24-.26.07-.3.17-.6.23-.9 0-.13.03-.33-.1-.4-.12.06-.3.09-.35.17-.2.3-.34.6-.53.92-.12.25-.17.25-.43.14a4.75 4.75 0 0 1-1.05-.65c-.25-.4-.52-.73-.65-1.08-.05-.09.1-.16.21-.16.28.22.62.5.87.74.17.19.22.37.33.61.03.06.3.05.35-.05.04-.1.08-.23.12-.43.1-.37.03-.88-.3-1.22-.42-.42-.7-.5-1.26-.9-.71-.48-.81-.45-1.07.03-.2.39-.35.82-.35 1.3.19 1.14.46 1.6 1.83 2.3a.38.38 0 0 1-.1.2c-.1.07-.15.11-.3.11-.49-.01-1.17-.02-1.7-.07-.15 0-.24-.06-.28-.14a6.98 6.98 0 0 1 .62-5.02.84.84 0 0 1 .33-.37c.29-.15.75-.34 1.03-.46.16-.08.25-.02.32.15.12.44.17.76.2 1.14.02.23.03.47.13.51.22.1.42-.09.53-.28.42-.68.53-1.44.22-2.26-.03-.07.04-.22.12-.27.95-.38 1.83-.77 2.8-1.14.53-.03 1.22.03 2.18.18.16.02.5.3.84.5-.44.22-1.08.69-1.53 1.11-.12.13-.28.3-.17.66.38-.02.56-.13.84-.2.43-.13.7-.2 1.04-.2a1.86 1.86 0 0 1 .28.04.51.51 0 0 1-.13.29c-.12.11-.31.1-.45.17-.14.1-.28.23-.19.4.07.11.25.24.38.22.45-.03.86-.08 1.32-.22.2-.17.99-1.2.54-1.29-.25-.05-.79-.09-1.02-.21a.47.47 0 0 1 .28-.32c.24.04.52.11.9.15a5.5 5.5 0 0 0 1.83-.08c.25-.1.33-.19.25-.35a2.88 2.88 0 0 0-1.7-1.48c-.34-.1-.75-.02-1.1-.02l-3-.03c-.06 0-.13-.06-.2-.1.06-.04.12-.1.17-.12.28-.17.6-.31.89-.5.42-.3.7-.71.81-1.25a.51.51 0 0 0 0-.28c-.28-1.03-.83-1.77-2-1.9l-.57-.1c-.22-.01-.3.08-.3.27 0 .37.02.56 0 .94 0 .18.1.25.3.27a7.17 7.17 0 0 1 1.59.46c.3.12.37.3 0 .62l-.84.6c-.15.1-.21.02-.23-.12-.02-.22-.02-.47-.06-.68a.47.47 0 0 0-.12-.21.61.61 0 0 0-.45.37c-.1.3-.18.57-.23.86a.91.91 0 0 1-.64.75c-.17.04-.32-.02-.29-.18a4 4 0 0 1 .12-.56c.27-.65.41-1.2.36-1.82a13.12 13.12 0 0 1-.1-1.29.38.38 0 0 1 .28-.4c.5-.2.98-.4 1.5-.56.19-.07.33 0 .47.18l1.07 1.25c.27.3.35.29.6-.05.21-.31.46-.6.81-.8.19-.09.19-.27-.03-.35-.4-.15-.8-.23-1.22-.37-.09-.03-.2-.03-.3-.07-.06-.03-.14-.11-.14-.17 0-.05.08-.14.14-.16.77-.27 1.4-.52 2.08-.79.19-.08.28-.23.25-.42-.02-.25-.1-.47-.13-.71-.04-.22-.19-.4-.42-.42-.93-.1-1.49.17-2.12.82a.58.58 0 0 0-.15.51c.11.02.2.02.29 0a3.8 3.8 0 0 0 1-.43c.16-.18.3-.37.39-.37.05 0 .15.23.18.32-.09.1-.18.22-.92.54a21.25 21.25 0 0 1-4.07 1.4c-.2.02-.18-.1-.13-.18.26-.46.55-.93.84-1.42.28-.17.36-.1.25.17a5 5 0 0 1-.3.58c-.03.18.05.23.16.21a1.56 1.56 0 0 0 1.16-1.39c0-.23-.11-.31-.36-.3-.35.02-.74.05-1.08.16-.66.16-1.05.62-1.26 1.21-.1.28-.15.56-.2.85-.06.3-.32.42-.6.57-.45.28-.68.22-1.07.3a6.42 6.42 0 0 0-.98.22c0-.11-.03-.21 0-.28.22-.35.42-.7.65-1.02.33-.48.24-.75-.28-.93a1.94 1.94 0 0 0-2.23.78c-.28.44-.48.98-.68 1.46-.12.29 0 .46.3.53.48.08 1.05-.1 1.48-.21-.03.14-.04.67-.09.79l-.27.6c-.18.3-.38.4-.7.24a14.7 14.7 0 0 1-1.46-.79 3.37 3.37 0 0 0-1.73-.56.84.84 0 0 1-.18-.01c-.2-.05-.22-.12-.1-.25.16-.15.17-.31.08-.5l-.64-1.2a3.26 3.26 0 0 1-.3-.7c.15-.03.74.17.74.28.04.65.24 1.12.58 1.53.26.26.5.1.53-.17.05-.34-.42-.9.19-2.19C26.7 1 27.5 0 28.02 0c.28 0 .51.12.68.53.1.19.16.63.24.82.08.19.17.28.37.26.5-.07 1.2-.37 1.77-.45.36-.05.64.85.82 1.18.28.48.41.9.73.38.54-.75.79-1.58 1.36-1.85a4.03 4.03 0 0 1 2.35-.4c.25.03.33.13.27.38-.05.15-.06.3-.1.46-.05.18.1.24.3.1.27-.16.47-.38.66-.62.48-.52.96-.78 1.91-.78.45 0 .8.1 1.2.23.3.11.35.16.26.47l-.27.9c-.09.4-.02.46.34.66.32.19.57.32.76.56a.65.65 0 0 0 .4-.14A4.81 4.81 0 0 1 43.8.74c.3-.16.38-.14.5.14.25.64.39 1.19.43 1.78.34-.39.73-.8 1-1.12C46.74.6 47.74.6 49.09.94c.33.1.07.43-.12.69l-.26.39c-.09.14-.15.15-.26.03-.17-.2-.34-.54-.5-.71-.31-.3-.85-.13-1.22.1a5 5 0 0 0-1.81 1.84 9.77 9.77 0 0 1-1.58 1.88c-.24.09-.24.08-.19-.17a6.65 6.65 0 0 0 1.1-3.45.83.83 0 0 0-.64.14c-.93.52-1.4 2.37-1.53 3.6-.03.27-.03.47-.26.6-.39.22-.54.32-.93.5-.1.01-.15 0-.23-.01.18-.82.05-1.4-.42-2.17-.08-.15-.19-.23-.36-.17-.1.06-.19.13-.4.3a3.5 3.5 0 0 0-1.1 2.69c0 .18.18.18.31.09.17-.14.26-.25.36-.43.23-.39.43-.8.65-1.19.04-.1.15-.16.23-.19.05.1.05.2.03.28l-.17.75a2.33 2.33 0 0 1-.74 1.02c-.07.1-.11.26.07.28.8.04 1.27-.05 1.96-.42.72-.37 1.2-.68 1.52-.64.22.03 1.3.34 1.4.42-.06.11-.14.15-.2.22-.12.1-.65.38-.88.58-.37.3-.5.45-.43 1.07.04.28.09.54.15.8.03.1.1.21.15.23.1 0 .19-.07.24-.14l.9-1.15a.7.7 0 0 1 .19-.11c0 .07.02.14 0 .18-.07.46-.3.8-.58 1.14l-.18.23c-.24.35-.26.4-.02.75.28.4.32.86.23 1.3-.06.32-.17.8-.23 1.28-.59.27-1.18.65-1.68.99-.31.28-.37.5-.37.57.15.18.29.21.58.33.4.13 1.17-.03 1.31-.46.05-.1.1-.3.1-.48.14.03.27.03.31.15.15.47.28.78.58 1.09-.04.12-.1.21-.23.4-.17.3-.32.55-.57.5-.1-.02-.2-.04-.31-.04-.19 0-.37 0-.35.33.04.43.1.62 0 1.06-.05.24-.13.5-.37.5-.58-.05-.97-.48-1.53-.47a2.2 2.2 0 0 1-1.07-.19c-.3-.15-.51-.32-.65-.4-2.36.38-4.96.86-6.97 1.3-.22.04-.2.18-.25.4-.1.52-.3 1.02-.39 1.53-.02.09.1.19.15.26.06-.05.14-.09.2-.16.12-.18.2-.39.34-.56.24-.26.88.02.99.37.1.21-.2.32-.19.56.02.2.32.39.32.65 0 .23-.1.38-.13.5.87.56 1.09 1.03 1.11 1.75.03.22.03.34.15.47.12.12.32.3.4.4.41.48.79 1.23 1.06 1.3.14.06.5.08.77.06l.05.14c-.09.22-.33.48-.56.55-.64.2-.95-.75-1.3-1.3a4.65 4.65 0 0 0-.54-.72c-.28.31-.74 1.55-.7 2.13.18.8.65 1.58.95 2.15.22.35.64.47 1.05.63.34.12.6.18.95.08.18-.04.25-.05.33-.12v-.72c.08-1.3.29-2.6.37-3.72 0-.3.23-.8.6-1.06.4-.28.57-.65.55-1.13a.88.88 0 0 1 .09-.46c.13-.26-.12-.26-.37-.3-.28-.04-.35-.15-.2-.4.25-.44.63-.8.92-1.23.19-.3.28-.5.35-.69-.08-.36-.05-.9.31-.92 1.19-.14 1.68-.03 2.43.53.05.08.06.19.06.28 0 .1.21.17.61.57.31.29.6 1.1.09 2.35.26.15.25.4.22.68-.02.24.06.27.28.36.28.1.43.26.66.84.2.55.81 2.01 1.16 5.6a14.92 14.92 0 0 1-.48 2.87c.24 1.72.51 3.84.73 6.73.17 2.74.22 4.14.27 5.27.49.04.64.14.74-.27l.1-.38c.09-.27.36-.18.37-.03l.3.87c.06.26.13.41.23.66.32.44.46.05.46-.27.04-.28.04-.52.04-.77.05-.32.22-.34.34-.12l.43 1.02c.16.35.33.44.4.23.05-.29.06-.8.05-1.21-.07-.23-.17-.49-.2-.69 0-.28.22-.22.27-.14a32.48 32.48 0 0 1 1.52 1.98c.28.17.45-.12.37-.23-.02-.11-.16-.25-.16-.47.04-.36.37-.2.45-.07.3.3.51.7.69 1.03a.65.65 0 0 0 .28.14c.06-.09.09-.23.08-.32-.02-.11-.15-.38-.19-.5.13-.21.4.06.48.14.1.1.19.25.3.45.17.26.45.48 1.07.57l.03.18c-.08.06-.22.12-.6.19a7.6 7.6 0 0 1-2.32.05 2.7 2.7 0 0 0-1.71.37c-.37.21-.74.18-1.18.17-.84-.04-1.68-.04-2.53-.06-1.77-.05-2.66.3-4.66.41-.65 0-2-.34-2-.75 0-.3.16-.53.16-.68 0-.19-.04-.24-.2-.28-1.43-.25-2.83.07-4.13.31-1.1.17-1.67 0-2.25-.46a2.91 2.91 0 0 1-.8-.85 11.37 11.37 0 0 0-2.22-2.11c-.93-.74-2.11-2.1-2.85-3.02-.17-.2-.46-1.15-.56-1.42a2.85 2.85 0 0 1-.13-1.23c.11-.1.25-.2.36-.21.62-.02 1.4-.32 1.42-.5.06-.14-.04-.26-.16-.23-.33.06-.77.06-1.1 0-.18-.03-.4-.1-.5-.2.12-.14.2-.24.32-.31a3.27 3.27 0 0 0 1.45-.56.35.35 0 0 0-.04-.18c-.4-.02-.72-.02-1.18-.13a1.05 1.05 0 0 1-.45-.26.72.72 0 0 1 .34-.28c.52-.17.9-.42 1.25-.75v-.2a7.1 7.1 0 0 1-1.42-.15c-.21-.11-.23-.13 0-.4.34-.08.96-.13 1.3-.24.24-.13.2-.35-.05-.42a7.63 7.63 0 0 1-1.21-.13c-.25-.23-.28-.25.04-.4l1.07-.22c.41-.14.35-.4 0-.5-.32-.04-.55-.11-.87-.17-.43-.21-.42-.23-.08-.37l.68-.19c.19-.04.38-.12.41-.26a.2.2 0 0 0-.02-.14c-.28-.08-.87-.02-1.13-.1-.1-.03-.19-.13-.24-.25l.25-.14c.32-.07.67-.21.92-.35a.81.81 0 0 0 .1-.24.6.6 0 0 0-.18-.09c-.34-.1-.53-.17-.87-.3-.19-.23-.18-.28.08-.38l.48-.13a3.1 3.1 0 0 0-.24-.48.48.48 0 0 0-.17-.17c-.28-.17-.68-.33-.98-.53-.56-.34-1.14-.68-1.58-1.12a3.63 3.63 0 0 1-1.17-1.91c-.09-.47-.1-.96-.01-1.42.05-.35.25-.72.44-1.06.07-.12.13-.23.05-.33-.28-.36-.51-.6-.86-.93a1.47 1.47 0 0 0-.55-.28c-1.11-.2-1.79-.42-2.3-.36a1.77 1.77 0 0 0-.42.64c.1.38.28.95.29 1.43.4.04.9.18 1.4.46.4.26.49.4.65.9.34 1.11.43 2.54.68 4.06.2 1.46-.55 2.66-2.02 2.71-.8.04-1.46-.12-1.67-.66-.26-.76-.46-1.4-.7-2.17-1.05.26-2.15.62-3.26 1.16-.98.46-1.78.9-2.81 1.33-1.97.8-4.14 1.3-6.08 1.3a8.6 8.6 0 0 1-2.96-.52c-.7-.28-.96-.57-1.12-1.46a8.48 8.48 0 0 1-.03-2.3c.09-.74.21-1.26.72-2.65.22-.62.51-.9.99-1.15a5.21 5.21 0 0 1 2.57-.5c.3.01.7.06.98.18.56.2.64.65.66 1.2 0 .33.02.66 0 .98-.03.26-.02.37-.02.47.7.1 2.18.06 2.88-.1.58-.11.81-.2 1.86-.7.73-.33 1.71-1.06 2.86-1.66zm22.28 12.45c.06.13.22.83.22 1.07 0 .57-.05 1.1-.11 1.98-.1 1.15.03 2.32-.13 3.46-.07.56-.2 1.42-.16 2.19 0 .3.06.39.23.42.56.13.93.19 1.21.19.86 0 1.42-.39 1.34-.87a11.63 11.63 0 0 1-.4-3.2l.23-.01c.31 1.3.58 1.94.67 2.62.12.61 0 1.67.13 2.2 1.43.16 3.14.02 3.66-.09.13-1.17.1-2.37-.42-2.82-.21-.2-.85-.68-1.13-.86a6.41 6.41 0 0 1-1.53-1.7c-.61-.96-.78-1.27-1.32-2.21-.5-.88-.84-1.63-1.18-2.14-.22-.35-.4-.56-.5-.65a6.58 6.58 0 0 0-1.3-.25c-.6-.06-1-.13-1.35-.5-.3-.3-.42-.55-.24-.95.1-.21-.03-.36-.18-.54-.18-.22-.13-.46.09-.63a3.55 3.55 0 0 1 1.72-.97l-.11-.48c-.28.04-.59.12-1.15.12-.77 0-1.47-.32-1.75-1.33-.1-.36-.55-1.39-.68-1.84-.1-.28 0-.63.14-1.03-.25-.3-.45-.63-.66-.91-.41-.18-.92-.35-1.25-.76-.14-.17-.2-.18-.33-.1-.11.09-.14.2-.3.35-.14.13-.45.14-.74.02-.28-.11-.42-.4-.41-.68.03-.37.21-.7.05-.85-.3-.3-.56-.65-.22-.9.1-.08.31-.06.44-.12a.93.93 0 0 0 .45-1.04c-.08-.31-.25-.45-.55-.48-.44-.03-.64-.25-.64-.54 0-.37.33-.53.63-.46.12.02.29.09.39.18.03.14.1.34.2.32.5-.06.9-.58.9-1.07 0-.17-.11-.25-.44-.18-.6.1-1.34.32-2.08.57-.4.14-.46.46-.24.84.08.08.13.18.19.28.1.14.05.26-.1.34l-.18.14c-.22.17-.25.35-.11.6.04.1.1.16.17.23.16.15.16.28 0 .46a.7.7 0 0 0-.14.76c.06.21.23 1.42.23 1.65-.01.63.11 1.25.07 1.84-.02.42.19.8.6 1.26.45.5.77.94 1.38 1.53.71.68 1.1 1.49 1.17 2.48.03.47.25.64.7.88.94.26 1.26.79 1.37 1.5.19 1.28-.43 2.52-1.58 2.4-.58-.05-.87-.4-.87-.71 0-.33.16-.65.43-.74.05-.02.22 0 .26 0a.56.56 0 0 1 .02.2c0 .08-.1.13-.15.3-.02.1-.02.33.09.47.12.16.33.22.56.14.26-.08.43-.3.55-.54.54-.96.27-2.1-.75-2.34-.32-.1-.79-.14-1.09.03-.63.4-1.17.62-1.7 1.1a.4.4 0 0 0-.13.3 52.05 52.05 0 0 0 .45 6.34c0 .45.33 1.02.46 1.1.37-.27.47-.54.8-1.06.15-.09.22-.03.27.1 0 .48-.1.94-.05 1.35.03.12.17.23.27.31.22-.2.23-.5.32-.74.25-.26.3.03.32.1-.02.3 0 .59.03.75 0 .12.12.14.22.17.27-.54.61-1.02 1.05-1.32.26-.19.43-.03.28.23a6.6 6.6 0 0 0-.6 1.58.93.93 0 0 0 .26-.03c.06-.02.14-.04.17-.19.19-.46.29-.52.5-.18.1.31.24.54.42.63.28-.08.3-.41.41-.65.3.02.32.34.47.5.1.08.95.11 1.1.1l-.02-.5c-.02-.99 0-1.78.06-2.76.08-1.02.19-2.07.24-2.97.04-.75.13-1.88.2-2.77zM51.26 15.9c-.23-.03-1.16-.09-1.4-.09-.21 0-.77 0-.79.29 0 .08.11.23.24.32.68.46 2.58.54 3.35.41.02-.18 0-.4-.06-.56-.19-.57-.89-1.07-1.53-1.28a5.26 5.26 0 0 0-2.06-.07l-.14-.35c.23-1 .2-1.93.06-2.67-.08-.4-.19-.6-.4-.93-.14-.18-.29-.18-.46-.04-.1.1-.22.22-.41.47-.65.78-.62 1.3-.5 2.48.07.05.15.1.37.1.2 0 .3-.09.32-.37.02-.39.11-.8.17-1.16.1.03.26.14.28.2.04.61-.16 1.24-.3 1.81.03.07.47.26.4.45-.05.14-.74.31-1.22.35a2.46 2.46 0 0 0-1.14-.94c-.28-.11-.59-.2-.36-.39l.38-.15c.65-.27.43-.64.4-1.27-.02-.24-.12-.31-.34-.3-.33.02-.54.05-.8.19-.5.3-.87.79-1.22 1.19-.06.06-.2.03-.32.03l.62-1.75c.06-.18.15-.37.35-.32 1 .23 2.03.04 2.74-.67.3-.32.3-.47-.09-.74a2.36 2.36 0 0 0-2.14-.46c-.3.12-.5.3-.55.61-.03.19.9-.07 1.22-.04.1 0 .16.09.23.14-.6.26-1.29.39-1.86.55-.19.05-.26 0-.3-.18a.81.81 0 0 1 .3-.82c.43-.37.78-.81.87-1.4.06-.31-.2-.79-.34-1.06-.02-.07.02-.13.02-.18.37.09.88.33 1.19.55a2.63 2.63 0 0 0 1.83 2.03c.33.11.66.24 1.01.34.32.1.56 0 .54-.34l-.01-.46c.27-.12.6-.22.67-.2.53.07 1.07.33 1.61.5.38.2.2.32-.07.28-.4-.06-.81-.18-1.2-.28-.25-.05-.46.06-.5.3.17.3.94.65 1.22.77.4.14 1.95.32 1.93-.1a5.45 5.45 0 0 0-.18-1.23c-.1-.34-.4-.54-.72-.62a8.38 8.38 0 0 0-2.07-.24c-.16 0-.46.1-.74.24l-.09-.46c-.14-.62-.44-.77-.9-.93.02-.18.02-.3.02-.5.41-.23.91-.62 1.1-.95-.16-.16-.46-.24-.62-.3-.63-.22-.7-.46-1.12-.65a2.8 2.8 0 0 0-1-.25c-.32 0-.75.08-.84.2-.05.08 0 .21.18.34.28.22.87.37 1.16.58.18.1.09.15-.03.17a2.88 2.88 0 0 1-1.14-.33c-.21-.11-.28-.2-.43-.26-.28-.1-.33-.02-.52.3-.14.24.03.36.23.47.67.37 1.16.59 2 .68.1.1.12.25.15.44-.84-.11-1.12-.12-1.64-.3-.83-.45-1.49-.8-2.29-1.07-.15-.06-.2-.16-.04-.26.84-.5 1.64-1.12 2.49-1.58.12-.07.31-.04.46-.01.41.07.55.2.96.32.24.1.48.16.52-.06.12-.61.32-1.07.5-1.33.37-.57.9-.87 1.6-1.3C50.33.9 51.18 0 51.49 0c.84 0 1.67 1.53 1.84 2.1.11.4-.56.76-.1.76.38 0 .74.06.74.4a2.9 2.9 0 0 1-.46 1.6c-.32.44-.65.9-1.08 1.19l.03.15c.72.54 1.51.93 1.51 1.75 0 .43-.12.74-.37.96-.2.16-.35.33-.33.58.03.57.06.89.03 1.44-.03.38-1.17.45-1.33.62a4.65 4.65 0 0 1-.31 2.83l.23.43c1.15.8 1.3 1.54 1.47 2.17-.32.31-.84.57-1.02.57-.75 0-2.4-.56-2.72-.6-.13.23-.06.44.11.9.15.43.43.72.43 1.1 0 .2-.03.49-.56.57a3.74 3.74 0 0 1-2.48-.68c-.15-.08-.28-.18-.43-.25-.25-.14-.57-.03-.65.25l-.17.56c-.1.31-.25.34-.67.25-1.18-.25-1.73-1.52-1.73-2.65.05-.23.24-.14.28-.03.07.55.31 1.08.49 1.53.22.43.5.66.97.72.1 0 .22-.18.33-.38.25-.53.49-.88.41-1.5-.03-.22-.09-.52-.25-.72-.07-.1-.2-.16-.33-.28.06-.1.16-.21.25-.18.52.18.91.18 1.24 0 .2-.1.3-.25.43-.42.45-.12 1.18-.28 1.8-.37.84 0 1.58.12 2.2.4zM19.4 36.32c.45 0 .99.22.98.7l-.1.25a.25.25 0 0 1-.1-.03c-.17-.36-.44-.57-.88-.48-.4.08-.76.29-1.02.58-.18.18.18.3.4.41l.47.28c0 .25-.63 2.01-.31 2.07.8.14 1.86-.96 1.97-2.51-.03-.56.29-.84.9-1.05.73-.27 1.07-.6.95-1.4-.06-.38.17-.27.38-.6l.24-.26c.04.11.12.38.1.49-.17.64-.04 1.14.28 1.72l.6 1.54c.13.15.26.03.4.19.51 1.03 1.09 2.58 1.87 3.24.64.53 1.32 1.46 1.93 2.02.35.31.6 1.1 1.04 1.25.32.12.3.28.26.57-.06.35.1.42.22.42.22-.03.42-.08.65-.1.48.1.37.44.01.5-.24.02-.47-.03-.71-.03-.09.02-.17.06-.25.1.02.09.06.18.13.24.39.21.67.24 1.05.32.07 0 .13.06.21.09-.03.06-.1.21-.18.25-.12.02-.28.02-.43 0-.35-.08-.66-.23-.99-.37-.11-.04-.18-.16-.29-.24l-.28-.08c0 .11-.03.21-.02.3.07.16.23.42.16.47-.02.05-.25.19-.31.15-.28-.13-.5-.45-.77-.6-.33-.2-.64-.2-.95.01-.33.23-.62.42-.93.6-.38.25-.79.65-1.24.65-.25 0-.88-.38-1.16-.46-.52-.17-.69-.19-1.24-.31-.47-.11-.88.06-1.35.22-.58.18-1.21.25-1.84.37-.35.06-.76 0-1.13-.02-.56-.02-1.1-.07-1.65-.06-.37 0-.76.07-1.12.19-.16.04-.34.07-.49.07-.25 0-.5-.06-.72-.22a1.06 1.06 0 0 0-.8-.15c-.26.05-.77.15-.98-.05-.26-.26-.89-.78-1.24-.71-1.1.18-1.64.81-2.72.53-.76-.2-1.26-.38-2-.13-.2.07-.68-.04-.87-.01-.28.03-.37-.02-.34-.3l.03-.15c.05-.29-.23-.5-.51-.4-.3.1-.49.41-.77.56-.53.26-1.03.18-1.6-.04-.4-.17-.8-.1-1.17.1-.24.12-.63.7-.9.7-.08 0-.22-.16-.27-.29v-.27a1.86 1.86 0 0 1 .61-.56 3.49 3.49 0 0 1 2.13-.28c.33.05.7.09 1.01-.04.25-.11.51-.13.74-.25.42-.23.79-.42 1.26-.25.35-.16.53-.27.7-.53.1-.42-.06-.86.05-1.27.09-.37.25-.84.37-1.2.07-.18.2-.16.32-.14.34.5.75.96.54 1.41-.14.32.59.54.47.86-.12.24-.07.4.18.42l.18.02c.22 0 .22-.14.22-.41 0-1.53-.49-2.59-1.01-3.79a1.77 1.77 0 0 1 0-1.4l.3-.94c.18-.45.55-.84.97-1 .05-.02.21.05.25.12.68.96.93 2.12.6 3.28-.1.34-.17.7-.24 1.08-.02.09.01.18.03.28.06-.06.15-.12.2-.19.34-.5.68-1 .95-1.54.26-.52.7-.73 1.23-.83.59-.12.75-.05.6.52-.09.33-.19.67-.3.97a1.93 1.93 0 0 1-.99 1.2c-.33.18-.64.43-.93.65-.11.15-.14.14-.02.29.11.14.28.05.42-.03.55-.4 1.12-.72 1.68-1.03.2-.09.27-.01.25.23-.03.28-.07.66-.35.76-.3.1-.71.15-.98.28a1.86 1.86 0 0 0-.47.4c-.18.28-.18.47-.33.76-.03.07 0 .18 0 .28.1-.03.2-.03.26-.07.36-.26.56-.56.92-.84.28-.21.34-.19.45.16.04.11.13.2.24.3.1-.07.23-.13.31-.25.33-.48.53-1.45.96-2.09.08 0 .17.06.23.08.09.22-.03.95-.01 1.11 0 .45.15.54.55.34.32-.16.6-.33.92-.5.33-.16.6-.03.4.23-.1.16-.25.28-.35.44-.05.1-.04.26-.05.38.09 0 .22.01.31-.03.25-.07.49-.18.73-.28.32-.15.66-.1 1.02 0 .35.12.7.2 1.03.3.27.06.56-.06.8-.17a8.64 8.64 0 0 1 1.48-.57 4.75 4.75 0 0 0 1.83-1.05c.12-.13.3-.2.45-.33.1-.09.2-.53.21-.67.1-.37-.06-1.24.02-1.63.05-.1.12-.3.18-.4l.16.04c.05.14.11.36.11.47-.02.67-.03 2.16-.1 2.8 0 .16-.27.4-.42.46-1.17.42-2.35.87-3.53 1.2-.53.15-1.12.07-1.68.08-.44 0-.88-.07-1.32-.04-.34.03-.84.35-1.19.4-.97.15-1.8.09-2.76.26-.89.17-1.74.41-2.58.65-.23.06-.35.13-.53.26-.07.05-.09.15-.11.22.09.05.18.11.27.1.22-.03.4.04.62.01.52-.08 1.04-.14 1.58-.19.11 0 .26.08.36.15a1 1 0 0 0 .82.17c.32-.07.6.01.88.21.48.4.6.41 1.16.1.43-.23.9-.23 1.38-.2 1.06.05 2.05.2 3 .38a2.05 2.05 0 0 0 1.98-.76c.27-.34.53-.68.74-1.01.18-.24.37-.4.65-.4.6 0 1.21-.14 1.8-.1.33.02.5.33.5.52-.03.2-.1.32-.33.31-.46-.03-.9-.05-1.34-.11-.47-.04-.84.1-1.12.45-.15.16-.11.28.11.33.1.01.22.06.32.07.24.02.35.14.42.36.04.14.33.28.82.3.62.04.94.04 1.48-.13.23-.08.38-.34.53-.53.36-.51.6-.65 1.23-.51.45.08.75-.4 1.15-.37.73.1-.1-1.38-.56-1.87-.64-.47-1.15-1.41-1.71-2.11-.41-.51-.9-.7-1.26-1.33-.68-1.13-1.24-2.5-2-3.4-.29-.31-.93-.1-1.25.07a.88.88 0 0 0-.37.26c.91.77.95 1.84.63 2.63-.09.25-.18.24-.44.14a9.6 9.6 0 0 1-.72-.29c-.16-.08-.3-.04-.46.06-.7.42-1.54.98-1.68.84-.28-.27-.73-1.11-.7-1.54.02-.35.14-.65.27-.85a6.98 6.98 0 0 1-.9-.57v-.31c.58-.56 1.36-1.1 1.96-1.1zm-17.42-8.1c.08.55.3 1 .84 1.2.76.32 1.77.5 2.62.52 1.02.03 2.13-.06 3.32-.28a21.59 21.59 0 0 0 5.36-1.87 14.83 14.83 0 0 1 3.25-1.24 7.37 7.37 0 0 1 3.6-.1c.54.13.65.13.57-.25-.03-.26-.1-.5-.15-.77-.13-.89-.09-1.73-.42-2.5a1.03 1.03 0 0 0-.5-.45 2.94 2.94 0 0 0-.85-.26 6.05 6.05 0 0 0-1.8 0 9.06 9.06 0 0 0-2.73.84c-.93.46-1.86.97-2.55 1.39a7.35 7.35 0 0 1-3.78 1.09c-.89 0-1.47-.03-2.36-.13a8.41 8.41 0 0 1-3.47-.95c-.24-.12-.37-.09-.46.16a7.79 7.79 0 0 0-.47 3.6zM43.38 38c-.3.04-.94.13-1.32.15a1.77 1.77 0 0 1-.36-.05c.06-.11.1-.15.16-.18.45-.23.94-.45 1.27-.7.19-.15.15-.28.1-.4-.63.27-1.21.4-1.87.5-.05 0-.08-.04-.1-.07.51-.38 1.2-.69 1.57-1.02.25-.2.06-.26-.02-.35-.54.19-1.08.37-1.65.52a.65.65 0 0 1-.26-.02c0-.11.11-.16.2-.22a8.68 8.68 0 0 0 1.59-1.04l-.03-.14-.19-.02c-.44.08-1.4.22-1.93.27-.26 0-.22-.15-.17-.18.42-.19.74-.41 1.05-.74.11-.12.1-.23.06-.29-.5.24-1.08.24-1.42.28-.06-.08.02-.18.06-.22.4-.16.5-.29.7-.43.14-.1.44-.3.5-.4.04-.05.03-.16 0-.2a6.51 6.51 0 0 1-1.24.3.62.62 0 0 1-.22-.03c.03-.1.05-.17.1-.2.59-.3.9-.42 1.3-.67a1 1 0 0 0 .28-.36l-.04-.06-1.42.44c-.1-.02-.14-.08-.22-.14a.57.57 0 0 1 .16-.17 7.6 7.6 0 0 0 2.64-1.52c.22-.18.42-.4.5-.56l-.06-.11a13.7 13.7 0 0 1-2.9 1.63c-.22-.02-.18-.15-.1-.2.1-.1.37-.41.6-.7a8.4 8.4 0 0 0 2.43-1.41c.06-.06.09-.22.05-.32-.48.21-1.2.62-1.83.86l-.12-.04c.02-.11.08-.22.14-.33a8.11 8.11 0 0 0 1.53-1.24c.02-.07.05-.15 0-.23-.4.22-.85.37-1.26.51-.14-.09.03-.33.05-.39.59-.48.87-.8 1.15-1.2.06-.1.06-.18 0-.33-.32.23-.79.37-1.15.51l-.1-.03c-.02-.12 0-.1.04-.18.42-.4.79-.7 1.05-1.15.02-.05.02-.17-.03-.25-.35.12-.62.22-1 .32a.3.3 0 0 1-.14-.03c.02-.12 0-.15.05-.24.26-.23.52-.53.76-.9.07-.09.06-.18.01-.25-.18.05-.52.1-.69.11-.07.02-.16-.06-.21-.15l.65-.46a.56.56 0 0 0 .11-.18c-.08-.04-.15-.08-.2-.06-.28.03-.54.06-.82.12-.61.13-.61.21-.67.39a13.61 13.61 0 0 1-.68 2.03c-.08.18-.19.36-.37.43-.26.14-.2.27-.28.53-.11.5-.19.46-.4.95a2.7 2.7 0 0 1-.57.94h-.14l-.05-.7c-.05-.47-.14-.37-.22-.84-.05-.42.06-.87.11-1.23.05-.42.1-1.14.1-1.55-.04-.03-.08-.1-.13-.13a.4.4 0 0 0-.13.18c-.09.32-.24.82-.28 1.15-.1.76.05 1.5 0 2.27a8.65 8.65 0 0 0-.1 2.4c.25-.03.24-.06.59-.2.46-.18.8-.25 1.27-.67.34-.28.28-.85.4-1.32.16-.56.38-.76.6-.85.07-.02.13 0 .2.02l-.33.82c-.11.3-.22.77.1.82l.23.04c.01.06.01.14-.03.19-.3.28-.54.63-1.91 1.02-.16.05-.5.46-.52.62-.03.22-.03.47.02.82.14 1.37.82 2.65 1.58 3.92.64 1.08 1.12 1.9 1.81 2.6.33.3.67.53 1.21.27.3-.13.66-.36.75-.44.04-.06.05-.15.04-.23zm-18.6-19.7c.03.54-.07.68-.76.5-.73-.22-1.06-.14-1.53-.67-.28-.33-.63-.33-.93-.02-.14.17-.24.33-.46.37-.04-.07-.09-.18-.06-.26.08-.32.27-.56.57-.7.21-.12.47-.2.7-.26.64-.25.95-.63.95-1.38 0-.43.14-.86.14-1.3 0-.32-.12-.43-.42-.32-.53.2-1.03.4-1.56.63a.64.64 0 0 0-.37.35c-.15.28-.24.56-.37.84-.12.16-.23.08-.26.02-.25-.4-.56-.82-.82-1.16-.23-.28-.23-.38.03-.57l.86-.76c.43-.4.6-.93.64-1.52v-.81c-.01-.35.06-.46.41-.41.25.03.5.1.71.2.4.17.78.4 1.17.6.2.13.36.3.38.57 0 .3-.2.5-.47.39a8.56 8.56 0 0 0-1.31-.46c-.23-.05-.45.07-.49.28-.09.51.35.54.31.68-.11.36-.7.47-.78.8-.07.3-.07.5.2.45.74-.15 1.46-.6 2.02-.84.49-.2.74-.13.8.41.26 1.58.5 2.94.68 4.35zm11.44-5.02c.14.04.32.12.43.22.27.63.84 1.32 1.57 1.7.23.12.46.16.65.18.07-.28.19-.74.19-1.25 0-.4-.12-.84-.5-1.21a4.84 4.84 0 0 0-1.11-.84.47.47 0 0 0-.55.05c-.01.09-.01.23.04.31.18.3.6.59.8.88.06.3.04.28-.2.18a4.7 4.7 0 0 1-1.33-1.16c-.21-.24-.35-.24-.5.02a4.5 4.5 0 0 0-.67 2 2.05 2.05 0 0 0 1.11 2.14c.35.15.6.11.77-.21.17-.33.3-.61.4-.97.06-.19.06-.26-.08-.43a6.15 6.15 0 0 0-.56-.62.81.81 0 0 0-.28-.23.37.37 0 0 0-.07.24c.01.28.06.58.07.87 0 .07-.01.19-.06.25a.57.57 0 0 1-.32-.37c-.18-.69-.15-1.3.2-1.75zM4.06 3.15l.54.64c.22.42.34.8.45 1.13.1.34.24.34.45.19a.76.76 0 0 0 .32-.65 2.17 2.17 0 0 0-.12-.59c-.2-.72-.06-1.34.57-1.8.22-.14.37-.39.3-.65-.1-.36-.19-.7-.33-1.04C6.2.23 5.98.2 5.84.27c-.7.37-1.54 1.44-.9 2.34.03.05 0 .22-.05.31-.6-.58-1.74-1.78-2.09-1.75-.1.17-.17.51-.1 1.21.1.68.27 1.38.78 2.33.13.21.25.34.44.5.12.1.37.22.56.22.06-.14.07-.39.03-.55-.17-.64-.34-1.12-.55-1.7zm47 2.68c.49-.1.89-.24 1.07-.31.8-.3 1.33-.82 1.5-1.7.07-.23-.07-.36-.33-.42l-.52-.14c-.5-.15-.61.06-.8.56-.16.37-.31.8-.71 1.11-.12.1-.54-.2-.66-.28a4.26 4.26 0 0 0 1.68-2.1c.1-.17 0-.34-.18-.4-.25-.07-.5-.1-.8-.07a4.01 4.01 0 0 0-2.88 2.14c-.04.1 0 .31.06.34.11.05.4.02.52-.03.07 0 .1-.11.15-.18.33-.58.88-.93 1.51-1.21.19-.1.25-.03.15.18-.12.16-1.26.96-1.28 1.07-.15.33.07.37.26.37l.34-.09c.12.6.48.99.9 1.16zm-17.22-.1c-.02.28-.12.44-.38.35-.52-.23-.83-.8-1.42-.99-.21-.07-.34.12-.37.32-.1.6-.19 1.12-.19 1.35 0 .82.44 1.26.88 1.7.13.11.13.19-.01.25l-1.8 1c-.17.12-.37.14-.56 0-.17-.15-.23-.35-.17-.54.21-.9.36-1.84.6-2.72.36-1.4.97-1.87 1.98-2 .37-.06.64-.06.92-.06.3 0 .42.06.46.33.04.18.05.39.05 1.01zm1.36 6.17c0-.68-.55-.87-.79-1.45-.22-.57-.42-.57-.88-.16-.52.5-.84.92-1.02 1.63a3.9 3.9 0 0 0 .01 1.96c.05.2.16.41.38.35a.37.37 0 0 0 .27-.37c0-.58.05-1.14.34-1.51.13-.17.19-.2.26-.25.1.14.06.3.06.33-.09.43-.15.87-.24 1.3-.06.28-.06.44.09.7.17.26.3.3.42.12.53-.85 1.1-2.07 1.1-2.65zM6.14 13.78a5.21 5.21 0 0 0-.37-1.67c-.22-.53-.56-.84-.78-.84-.2 0-.68.56-.84 1.01-.4 1.15-.19 2.08.48 3.18.09.14.27.4.53.14.2-.24.36-.4.55-.65.26-.31.44-.75.43-1.17zm25.07 17.96c-.12-.96-.84-1.89-2.05-2.7-.25-.1-.27.07-.27.28.12 1.33.27 2.67.44 3.86.05.26.17.33.41.2.36-.17.7-.34 1.05-.54.32-.19.52-.36.42-1.1zm10.96-10.89c-.2-.46-.5-.66-.97-.7-.47-.05-.86-.11-1.35-.19-.28.04-.47.26-.63.63.01.08.04.17.1.19.16.06.44.11.48.14.06.04.24.23.28.38-.04.13-.2.27-.42.3-.28.06-.28.29-.23.58.04.28.14.5.36.8.19-.1.28-.15.42-.3a.63.63 0 0 0 .19-.54c-.04-.6.3-1.04.71-1.44l.1.05c-.07.35-.25.6-.3.94-.08.7-.32 1.3-1.02 1.78.22.28.75.07.92-.1.29-.26.52-.58.63-.98l.11-.35a.28.28 0 0 1 .1.02c.13.26-.06.8.03 1.18a.87.87 0 0 0 .35-.28c.26-.62.39-1.53.14-2.1zm8.76-14.33c-.12 0-.6.01-.74.04-.11.03-.17.29-.11.34.13.09.22.16.33.2.44.05.63.05 1.14.13.08.01.18.06.25.14-.05.04-.1.1-.19.12-.93.11-1.2.1-1.86-.34-.18-.08-.24.05-.19.22.1.32.36.58.84.77.38.17.75.37 1.15.37.6-.02 1.02-.06 1.51-.14.26-.04.52-.19.56-.32.05-.13 0-.42-.15-.6l-.35-.36c-.56-.59-.92-.96-1.75-.87-.11 0-.28.3-.44.3zm.58 6.2c0-.34-.05-.6-.14-.9a1.58 1.58 0 0 0-1.38-1.14c-.34 0-.51.06-.37.38.04.1.34.43.46.7.16.31.31.64.31.91a.93.93 0 0 1-.08.38c-.2-.11-.32-.37-.37-.49-.09-.3-.1-.47-.2-.64-.08-.15-.19-.19-.23-.19-.13 0-.19.05-.2.23a2.67 2.67 0 0 0 1.46 2.67c.26.1.33.03.41-.2a5.58 5.58 0 0 0 .33-1.7zM5.58 9.12c-.13.5-.15.83-.14 1.1.01.36.19.66.52.8a6.23 6.23 0 0 0 3.1.5c.03 0 .13-.07.13-.1-.02-.22-.14-.25-.33-.3-.44-.16-1.61-.47-1.77-.66-.12-.17-.06-.34.13-.37a1 1 0 0 1 .51.08c.3.14.54.33.84.5.1.05.2.15.32 0 .08-.14.16-.27-.02-.4l-.69-.55c-.6-.47-1.49-.7-2.6-.6zm21.29 7.71c-.05.1-.14.2-.25.33-.1-.11-.22-.19-.27-.3a22.36 22.36 0 0 1-1.12-3.73c-.02-.24.07-.33.28-.32.53.06 1 .3 1.46.54.07.06.13.22.14.33.1 1.05-.04 2.23-.24 3.15zM3.95 18.97l-.08-.03c.05-.64.37-1.54.25-1.77a1.3 1.3 0 0 0-.64.06c-.22.1-.32.32-.36.7-.06.82.07 1.67.25 2.41.05.28.17.31.37.11.53-.48 1.03-1.1 1.43-1.57.28-.36.33-.7.22-1.05-.1-.29-.29-.4-.5-.4-.24 0-.32.18-.38.35-.13.46-.34.79-.56 1.19zm42.97-1.9c.42 0 .76.08.88.12.5.18.72.4.81.58.1.17.1.26.1.34a.9.9 0 0 1-.83-.17c-.45-.42-1.1-.6-1.49-.42.03.3.3.65.54.84.78.6 1.47.7 2.3.51.15-.05.2-.25.2-.42-.03-.54 0-1.11-.5-1.58-.57-.53-1.37-.51-1.92-.28-.09.05-.23.11-.3.19.02.06.09.2.21.28zM27.56 2.32c.17-.04.1.03.05.18-.1.48-.37.85-.7 1.19-.2.16-.23.42-.1.55.24.23.52.03.7-.15.38-.41.61-.93.83-1.43.26-.62.2-1.4-.1-1.92-.15-.22-.4-.25-.58-.08a3.23 3.23 0 0 0-1.14 2.23c0 .13.01.35.1.4.42-.2.68-.66.93-.96zm11.36 22.52c.04-.04.12 0 .15.02 0 .5.12 1.1.18 1.28a.58.58 0 0 0 .19-.13l.14-.68c.31-.96.08-1.72-.32-2.6-.07-.16-.18-.18-.35-.1-.12.05-.23.13-.12.3.09.14.08.25 0 .4-.13.24-.23.51-.37.81-.05.19-.06.33-.08.56l-.14 3.44c.06.19.17.15.23.04.08-.19.13-.36.17-.56.1-.54.25-.56.28-1.1.05-.56.05-1.11.03-1.67zM7.39 3.57c-.01-.05 0-.18.02-.21l.57.1c.28.05.42.01.52-.2.15-.41.07-.78-.36-.87-.65-.12-1.3.1-1.86.22-.36.13-.45.5-.18.77.41.4.81.8 1.25 1.15.36.29.77.36 1.18.28.19-.05.35-.17.35-.33a.47.47 0 0 0-.14-.34 1.26 1.26 0 0 0-.33-.18zM6.8 28.63c-.03.19-.12.35-.43.33a3.83 3.83 0 0 1-.83-.12c-.3-.08-.38-.2-.4-.3-.1-.57-.04-1.33.09-1.85.04-.1.35-.28.48-.35.4.04.72.1 1 .18.12.07.34.4.36.45.06.18.06.37.06.56-.11.46-.25.86-.33 1.1zm31.08 16.8c.75.28 1.33.36 1.55.32a.45.45 0 0 0 .34-.5c0-.2-.1-.38-.37-.45-.57-.12-1.15-.2-1.7-.4-.3-.1-.55-.09-.83.13-.43.32-.84.5-1.37.6a.65.65 0 0 0-.34.12c-.15.2-.11.33.12.46.92.08 1.79-.1 2.6-.28zm3.07 1.83c.46 0 .8-.1 1.12-.3a2.8 2.8 0 0 1 1.35-.6c.2-.04.41-.25.43-.42 0-.23-.19-.26-.34-.31-.45-.14-.98-.12-1.44-.24-.15-.02-.36-.02-.48.07-.96.97-1.15.74-1.6 1.22-.12.16-.12.35.21.44.26.06.6.12.75.14zM25 10.86c0-.23.43-.6.67-.57.04.02.09.05.13.1l.09.5c.03.2.1.37.2.37.48.04 1.05-.3 1.32-.72.28-.44.49-.95.73-1.44.22-.53.57-.2.45.12-.31.92-.7 1.86-.99 2.77-.1.33-.14.35-.46.24l-1.54-.55a.8.8 0 0 1-.6-.81zM36.8 33.85c.25.12.23.07.56-.12a.76.76 0 0 0 .3-.54l-.1-2.01c0-.07-.09-.17-.12-.17-.28 0-.62-.06-.84.06-.43.24-.8.57-1.18.88-.15.15-.01.3.21.28.19-.08.4-.17.59-.27l.35-.19c0 .07 0 .12-.03.17-.34.37-.6.82-1.05 1.06-.12.02-.09.32.09.33.14-.02.42-.17.5-.25.29-.22.4-.36.69-.6.1.15.16.26.02.38l-.75.67c-.11.11-.03.22.1.2.2-.02.38-.1.53-.2.2-.12.35-.29.52-.47 0 .02.1.13.09.13-.14.22-.3.42-.48.66zm-23.82-23.4c-.11 1.12-.39 2.25-.9 2.77-.4-.79-.68-1.7-.96-2.5.22-.28.5-.7.62-1.02.06-.16.15-.19.3-.08.33.31.62.56.94.82zm1.33 16.04a.93.93 0 0 1-.5.64c-.2.1-.52.23-.88.23-.17 0-.3-.04-.32-.19l-.3-1.5c0-.12.1-.28.2-.38.32-.3.74-.39 1.18-.41.1 0 .4.13.45.2.09.47.14.9.17 1.4zm-4.06-6.03c-.6.1-1.12.63-1.03.84.01.47.1.91.36 1.31.15.24.27.3.52.13a4.47 4.47 0 0 0 1.78-1.82c.12-.22-.12-.29-.34-.4-.3-.14-.61-.14-.78-.06-.17.09-.23.27-.26.4-.19.68-.23.7-.43.94a1.8 1.8 0 0 1 .18-1.34zM14.28 46c.03-.08.05-.13.12-.14.1-.02.37-.02.4 0 .53.15 1.02-.23 1.51-.3.11-.02.24-.04.37-.02.32.01.65.02 1.01-.07.15-.06.34.02.51.1.57.24 1.12.38 1.67.06.29-.2.58-.38.88-.53.51-.2.56.14.58.33a.48.48 0 0 1-.28.29c-1.77.63-3.3.64-5.28.4a4.47 4.47 0 0 1-1.5-.12zM3.2 27.39c-.04.24.04.6.23.8l.09.08-.05.15a4.19 4.19 0 0 1-.91-.12c-.13-.05.23-.24.26-.47.06-.45.19-1.2.19-1.75 0-.1-.24-.2-.24-.28a.34.34 0 0 1 .11-.19c.66.16.75.15.88.46.13.28.25.74.45 1.08.03-.02.09-.04.12-.07.03-.3.02-.75.03-.96.01-.17.09-.25.27-.25.13 0 .48.01.62.04.02.03.02.05.02.1 0 .08-.11.11-.14.12-.31.11-.33.14-.47 1l-.15.98c-.02.08-.03.23-.08.33-.08 0-.23 0-.26-.08-.21-.52-.5-1.08-.76-1.55h-.1c-.07.2-.07.4-.12.59zm20.47 4.59c.03-.2.03-.4 0-.57-.1-.32-.25-.63-.2-.9.14-.49.24-1 .2-1.5a6.27 6.27 0 0 1 .05-1.86c.17.06.3.17.34.28.04.2.09.4.09.57.06.53.02 1.04-.01 1.56-.08 1-.17 1.93-.14 2.92.02.15.05.34.1.48.12.3.06.63-.18.97l-.15-.07c-.11-.41-.4-.98-.51-1.43-.02-.08 0-.15.1-.17.2-.03.31-.15.32-.28zm9.68-7.4c-.4-.04-.83-.4-.8-1.05-.2.27-.45.55-.68.66.02.22.57.54.95.7.22.28.37.42.65.74.17-.43.38-.8.44-1.03a3.3 3.3 0 0 0-.32-1.51c-.31-.35-.8-.63-1.27-.94-.23.25.06.79.64 1.11-.13.36-.11.9.46 1.21zm-27.82-1c.29-.02.6-.06.94-.11.22-.06.47-.11.48-.4.02-.24-.28-.38-.48-.44a4.43 4.43 0 0 0-2.93.3c-.25.13-.46.27-.6.42-.22.24-.18.39 0 .52.18.14.45.29.8.37.04-.2.08-.54.15-.74.32-.52 1.21-.62 1.83-.64.15.03.15.22.03.26-.18.14-.53.1-.62.26 0 .15.15.21.4.2zm1.99 5.4c-.01-.11.12-.25.13-.33.03-.15.03-.35.03-.5.02-.4-.07-1.06-.07-1.45l.22-.14c.07-.03.13.02.21.1.19.18.6.64.8.82.02 0 .05-.05.06-.06v-.3c0-.2-.04-.35-.1-.57-.07-.25-.25-.43-.05-.45l.76-.05c-.07.15-.13.26-.16.48-.03.3-.1.59-.1.88.04.44.13.89.17 1.26 0 .03-.06.07-.11.1a.7.7 0 0 1-.22-.14c-.37-.38-.63-.72-1.11-1.18a.4.4 0 0 0-.07.2c.1.3.21.62.41.9.12.16.15.21-.16.33-.17.05-.45.13-.63.09zm13.6.08c.44-.27.62-1.17.5-1.6a.58.58 0 0 0-.14-.31c-.14-.14-.3-.1-.37.14-.09.35-.13.72-.17 1.05a3.26 3.26 0 0 0-.99-.03c-.33.05-.98.24-.98.64 0 .65 1.5.5 2.15.11zm2.63-4.5a.74.74 0 0 0 .18.52c.67.7 1.6 1.42 2.18 1.56.52.14.96.45 1.28.87l.3.42c.23.28.52.46.9.52.07.02.16 0 .23 0-.06-.23-.06-.26-.22-.43-1.38-1.23-2.84-2.19-4.26-3.18-.14-.09-.4-.25-.6-.28zM9.36 40.38v-.02c-.14-.59-.3-1.15-.44-1.72-.12-.2-.32-.06-.3.03-.15.41-.3.83-.43 1.27a.61.61 0 0 0 0 .4c.22.6.43 1.21.67 1.83.04.08.15.13.21.18.06-.07.14-.16.16-.23zm2.44-14.6c.04.12.1.26.11.39.01.1.01.17-.02.22 0 .03-.06.02-.12 0-.29-.14-.61-.32-.87-.15-.1.07-.06.24.06.3l.82.33c.37.14.47.3.4.74-.07.42-.38.59-.74.69-.15.04-.31.06-.51.04-.03.13 0 .28-.09.33-.06.04-.08.03-.15.06a3.6 3.6 0 0 1-.28-1.02c0-.08.03-.09.06-.12.12.04.23.2.42.38.22.08.6.14.66.08.11-.14.16-.23.16-.28 0-.1-.13-.19-.25-.23-.28-.16-.44-.22-.74-.34-.28-.11-.33-.36-.33-.6 0-.25.14-.52.54-.65l.46-.13c.04-.08.07-.36.07-.41.06-.01.18-.03.2.01zM6 8.26c-.57 0-.92-.14-1.48-.09-.4.03-.79.2-1.18.35-.23.12-.19.22.02.3.2.1.34.15.54.14l3.6-.33c.07-.1.12-.17.19-.32-.44-.29-1-.23-1.54-.39zm13.24 17.23c-.05.03-.13.1-.27.24a.73.73 0 0 1-.23.17l-.16-.1.26-.6c.09-.23.18-.26.4-.11l.2.17c.18.14.4.25.6.16.2-.1.1-.38-.1-.48-.3-.19-.53-.24-.9-.46-.21-.13-.3-.33-.3-.61 0-.23.15-.51.45-.54.11-.03.25-.03.4-.03.2 0 .33-.03.45-.2.1-.11.2-.19.23-.06.07.34.03.7-.06.91-.06-.01-.1-.01-.17-.04-.17-.28-.4-.41-.6-.38-.19.02-.28.16-.14.31.25.26.6.4.9.57.43.23.4.56.35.8a.6.6 0 0 1-.63.5c-.22 0-.43-.06-.68-.22zm-.84-.86c.04.72-.54.98-.95.97-.38 0-.6-.08-.74-.46-.06-.23-.14-.9-.18-1.16-.08-.1-.2-.13-.4-.18-.1-.02-.1-.1-.12-.18.14-.07.39-.14.56-.17.41-.06.41-.04.48.34.05.42.12.87.2 1.3.14.42.59.33.59-.03 0-.33.02-.71 0-1.07-.07-.08-.19-.26-.32-.38-.16-.19-.16-.24.06-.33.52-.2.56-.33.63.36.03.26.19.66.19 1zm.48 3.67l1.3-.3c.35-.1.54-.47.52-.8-.04-.25-.19-.39-.45-.44a3.54 3.54 0 0 0-1.68.23c-.04.44.12.93.31 1.3zM6.91 23.66c-.16 0-.39.04-.55.09-.33.08-.5.11-.82.08-.19-.02-.49-.14-.66-.18-.2.01-.4.2-.47.47-.04.28.13.48.44.53.62.14 1.2.25 1.79.25.22 0 .25-.03.26-.21.02-.35.01-.65.01-1.03zm14.9-13.3a3.77 3.77 0 0 1-.82-.25c-.09-.03-.15-.26-.15-.38.02-.32-.06-.65 0-1 .02-.18.13-.23.3-.2l.73.19c.32.13.79 1 .65 1.31-.14.28-.39.32-.7.34zM16.42 25.6a.93.93 0 0 1-.26.25c-.3.2-.86.33-1.24.58-.19.13-.1-.32-.11-.49l-.2-1.04a1.36 1.36 0 0 0-.34-.56c-.1-.13.2-.18.37-.23l.6-.19c.23-.07.28.09.22.14-.31.22-.28.32-.25.6.05.29.05.6.06.83.05.11.11.26.2.36.18-.05.39-.2.39-.25l.04-.6c.23-.27.26-.14.3-.01.1.22.14.36.22.61zm.93-6.6c-.13-.4-.26-.83-.36-1.24-.03-.17.11-.27.27-.24.56.12 1.1.25 1.67.4.18.04.18.19.06.31-.36.37-.74.69-1.24.91-.26.14-.26.04-.4-.15zm24.1-14.77a4.1 4.1 0 0 0-.63-1.54c-.26-.32-.56-.35-.84-.06-.1.1-.18.19-.24.28-.28.36-.26.47.16.67.64.19 1 .56 1.45 1.03zM11.9 40.4a2.6 2.6 0 0 0-.83.74l-.58 1c-.07.16-.18.55-.22.8l.1.1a4.65 4.65 0 0 0 .83-.6c.47-.46.58-.9.84-1.83l.05-.11zM52.9 2.13c-.25-.52-.71-1.22-.9-1.47-.26-.32-.38-.34-.66-.06l-.25.28c-.28.35-.45.78 0 .84.4.07.93 0 1.18.18l.54.36zm-47.18 25c-.07.69-.05.69-.05 1.16 0 .26.12.26.35.28.21 0 .28-.19.33-.44l.14-.94c.03-.3.05-.48-.13-.51-.35-.05-.6.17-.64.46zm3.63-12.58a3.6 3.6 0 0 1-.87-.3c-.47-.24-.91-.75-.84-1.21.04-.22.16-.28.34-.18.36.26.66.65.92.93.21.22.37.47.46.76zm3.5 11.06c0 .19.16.74.22 1.05.02.17.06.3.3.27.32-.05.43-.18.38-.44-.03-.28-.1-.57-.16-.87-.06-.28-.18-.44-.43-.43-.26.03-.35.12-.32.42zM26.66 8.05c-.2.35-.37.7-.6 1.01-.18.26-.46.41-.71.63-.24.06-.25.04-.17-.2.36-.6.7-1.06 1.1-1.57.28-.15.38-.03.37.13zm18.67 9.66c0 .22-.04.43-.11.64l-.06.03c-.29-.73-.99-1.9-.94-2.02-.02-.08.17-.13.25-.11l.32.13c.18.43.37.84.54 1.33zm2.03-15.42c.03 0 .15.04.23.11.14.11.16.24-.03.32-.44.2-1.17.6-1.65.77-.09.05-.2.06-.15-.11.09-.28 1.2-1.1 1.6-1.1zM31.44 22.5c-.3 0-.56.24-.56.48 0 .27.35.56.66.56.28 0 .47-.18.47-.46.02-.34-.22-.58-.57-.58zM6.91 46c-.06 0-.18.2-.28.21-.38.03-.5.07-.5.23 0 .21.1.32.28.27.56-.17.69-.15 1.42-.1l.25-.3c-.09-.07-.15-.16-.25-.19A4.16 4.16 0 0 0 6.9 46zM39.7 19.21c-.12.07-.22.15-.19.29l.1.22c.21-.04.45-.06.68-.04.3.1.8.2 1.11.26.08.02.17-.06.23-.14l-.07-.12c-.43-.43-1.14-.44-1.86-.46zM19.19 9.8l-.76-.34c-.05-.02-.08-.16-.03-.22.17-.25.34-.48.52-.7.1-.1.17-.09.21.03l.38.95c.08.19 0 .26-.32.28zM9.18 19.15c.18-.04.27.08.25.24-.02.1-.05.13-.2.2-.58.19-1.1.18-1.7.06-.18-.02-.3-.06-.33-.1a.35.35 0 0 1 .22-.16zm5.34-14.93c-.06.01-.12.06-.16.03-.5-.28-1.2-.65-1.65-.94-.1-.15-.04-.2.1-.18.37.12 1 .3 1.36.44.24.12.31.33.35.65zm7.58 10.86c.06.09.11.17-.04.42-.26.4-.32.8-.49 1.33-.04.12-.11.24-.17.32-.05.06-.14.08-.2.11-.02-.06-.06-.15-.05-.2l.5-1.64c.08-.2.23-.3.45-.34zM40.28 9.6a9.03 9.03 0 0 1-1.96-.68c-.21-.19-.17-.18.1-.19l1.84.5c.35.26.28.25.02.37zM32.7 21.25a.62.62 0 0 0-.28-.35c-.1-.05-.27-.03-.37.06-.1.06-.22.16-.3.26-.14.2-.05.39.17.55.18.13.39.12.51-.06.11-.14.21-.24.27-.46zm-12.26-9c-.07.15-.09.29-.18.39-.25.3-.53.58-.8.87-.27.12-.28-.04-.25-.17.25-.47.6-.86.93-1.24.03-.04.14-.07.18-.04.05.03.08.1.12.19zM9.93 9.7c-.07.04-1.05-1.5-1.17-1.8 0-.06.16-.18.24-.24.06.04.14.1.17.15.31.67.86 1.8.82 1.85zm-.42 6.13l-.42-.19c-.3.07-.56.2-.9.25-.36.02-.51-.05-.65-.14.05-.12.25-.23.54-.28.17-.03.9-.1 1.54-.04zM47.4 8l1.08.81c.1.08.09.23.13.32l-.24-.03c-.46-.23-.95-.42-1.25-.87-.26-.39.16-.34.28-.23zM20.04 9.34l-.12.19c-.04-.06-.1-.11-.14-.19-.09-.18-.09-.4-.1-.6-.06-.46-.04-.7.24-.99.12-.13.2-.05.2.07.03.58 0 .95-.08 1.52zm-6.92-1.7c.1.05.28 1.15.4 1.72.03.13.02.3 0 .42-.05.13-.1.16-.2.05-.19-.59-.28-1.18-.38-1.74-.03-.17-.03-.35.18-.44zm31.57 6.74c.44.25.96.7 1.18.96.05.05.03.2 0 .31-.08-.03-.16-.03-.22-.08-.37-.23-.71-.5-1.06-.74-.04-.03-.05-.3-.09-.36a.4.4 0 0 1 .19-.1zM24.34 4.79c.07.05.1.1.15.16a.4.4 0 0 1-.15.17c-.32.12-1.06.12-1.36-.06a.95.95 0 0 1-.16-.18l.21-.12c.44-.1.93-.06 1.3.04zm8.37 1.79c.3.46.43.91.56 1.41 0 .03-.02.04-.02.12-.07-.05-.11-.06-.14-.1a2.6 2.6 0 0 1-.74-1.23c-.04-.1-.02-.2.12-.28a.5.5 0 0 1 .22.08zm-5.79 31.85c.36-.2.46-.36.77-.55.13-.08.12-.1.05-.22-.07-.14-.14-.14-.32-.1-.35.13-.49.21-.84.35-.15.06-.2.12-.05.24zm-10.9-33c.43-.3.9-.45 1.37-.63.08-.02.19.01.27.09l-.16.14c-.38.27-.75.47-1.1.68-.11 0-.28-.03-.38-.06a.28.28 0 0 1 0-.22zM4.9 14.55c-.1-.5-.24-1.11-.15-1.56.02-.1.09-.19.2-.28.08.09.18.15.19.23.09.53-.06 1.12-.1 1.6zM29.13 3.32c.32-.26.86-.54 1.25-.73a4.7 4.7 0 0 1-1.28 1.37c-.03-.2-.1-.53.03-.64zm-5.96 19.44c0 .2.1.43.26.59.1-.08.17-.22.17-.28.04-.4.06-.8.06-1.2.05-.16-.18-.2-.24-.04-.09.3-.24.62-.25.93zm20.44 16.1l-1.72.58c.4.45 1.56.08 1.72-.58zM18.26 7.6c.03.46-.06 1.03-.3 1.45-.04.05-.1.09-.18.14-.02-.08-.1-.14-.09-.2.1-.47.21-1.02.33-1.49.01-.03.09-.05.13-.06zm24.1 32.9c.37 0 .73 0 1-.08.26-.16.25-.49.33-.71l-.1-.03-.4.37a5.8 5.8 0 0 1-.84.17c-.05.06-.1.1-.13.17.05.06.1.11.15.12zm-3.44-18.43c.12-.27.23-.52.32-.77.02-.05-.06-.2-.16-.29a.73.73 0 0 0-.2.1 3.72 3.72 0 0 0-.55.82c.04.06.22.06.58.14zm-28.81-7.3l.36-.92c.06-.02.12-.02.17 0 .05.29 0 .73-.02.91-.06.25-.15.4-.37.51a1.52 1.52 0 0 1-.14-.5zm21 6.84c-.25 0-.45.15-.47.37 0 .19.2.36.43.36.24-.01.41-.18.43-.39 0-.22-.14-.34-.39-.34zm-.22 2.29c0-.25-.15-.42-.42-.42-.22 0-.37.14-.37.33 0 .23.23.45.46.43.2 0 .33-.15.33-.34zm-.94-2.35c0 .2-.13.33-.33.33-.26 0-.46-.16-.46-.4a.4.4 0 0 1 .38-.37c.2-.02.4.16.4.44zm.16.93c-.22 0-.37.14-.37.33 0 .2.2.37.42.37.17 0 .33-.18.33-.37 0-.18-.16-.33-.38-.33zm-2.44 16.64c-.3 0-.54.14-.53.3.01.16.24.37.41.37.2 0 .3-.2.3-.44-.02-.19 0-.23-.18-.23zM9.93 3.9a3.23 3.23 0 0 1 .06-1.63c.24.48.26 1.27-.06 1.63zm20.94 16.3c-.15.1-.28.25-.26.31.02.11.14.21.21.33a.6.6 0 0 0 .25-.3c.03-.07-.1-.19-.21-.34z" fill="#53565a"></path></svg><svg xmlns="http://www.w3.org/2000/svg" role="img" version="1.1" height="30" width="138" viewBox="0 0 138 30" class="gh-wordmark" focusable="false" aria-hidden="true"><title id="gh-wm-scopus-preview">Scopus Preview</title><g fill="#f36d21"><path class="a" d="M4.23,21a9.79,9.79,0,0,1-4.06-.83l.29-2.08a7.17,7.17,0,0,0,3.72,1.09c2.13,0,3-1.22,3-2.39C7.22,13.85.3,13.43.3,9c0-2.37,1.56-4.29,5.2-4.29a9.12,9.12,0,0,1,3.77.75l-.1,2.08a7.58,7.58,0,0,0-3.67-1c-2.24,0-2.91,1.22-2.91,2.39,0,3,6.92,3.61,6.92,7.8C9.5,19.1,7.58,21,4.23,21Z"></path><path class="a" d="M20.66,20A6.83,6.83,0,0,1,16.76,21c-3,0-5.23-2.18-5.23-6.29,0-4.29,2.91-6.11,5.28-6.11,2.16,0,3.67.54,3.85,2.11,0,.23,0,.57,0,.86H18.81c0-1-.55-1.25-1.9-1.25a2.85,2.85,0,0,0-1.35.21c-.21.13-1.85.94-1.85,4.11s1.9,4.65,3.59,4.65a5.91,5.91,0,0,0,3.2-1.2Z"></path><path class="a" d="M27.29,21c-3.28,0-5.46-2.44-5.46-6.19,0-3.46,2-6.21,5.75-6.21,3.3,0,5.49,2.37,5.49,6.21C33.06,18.5,30.85,21,27.29,21Zm0-10.69a3.3,3.3,0,0,0-2,.65A5.83,5.83,0,0,0,24,14.73c0,3.74,2,4.6,3.56,4.6a3.45,3.45,0,0,0,2-.65A5.53,5.53,0,0,0,30.9,15C30.9,12.86,30.2,10.36,27.31,10.36Z"></path><path class="a" d="M40.37,21a5.63,5.63,0,0,1-2.6-.57v5.46h-2V12.23c0-.91-.05-1.72-.1-2.31l-.1-1H37.4l.31,1.74a4.86,4.86,0,0,1,4-2.05c2.39,0,4.26,1.56,4.26,5.72S43.69,21,40.37,21Zm.91-10.61a4.49,4.49,0,0,0-1.56.31,11.57,11.57,0,0,0-2,2.11v5.8a4.35,4.35,0,0,0,.7.34,4.12,4.12,0,0,0,1.61.34c2.57,0,3.74-1.9,3.74-4.73C43.82,12.94,43.51,10.44,41.27,10.44Z"></path><path class="a" d="M58.36,20.74H56.54L56.22,19a4.06,4.06,0,0,1-3.85,2.05c-2.08,0-3.77-.86-3.77-3.87V9h2V15.8c0,1.92.16,3.54,2,3.54a4.47,4.47,0,0,0,2-.47,6.77,6.77,0,0,0,1.64-2.08V9h2v8.53a19,19,0,0,0,.1,2.31Z"></path><path class="a" d="M64.86,21.07a6.87,6.87,0,0,1-3.67-1l.23-1.87a5.54,5.54,0,0,0,3.28,1.2c1.66,0,2.44-.75,2.44-1.66,0-2.39-5.88-2.26-5.88-5.9,0-1.77,1.38-3.22,4.21-3.22a6.59,6.59,0,0,1,3.38.88l-.21,1.87a4.67,4.67,0,0,0-3.15-1.14c-1.33,0-2.24.52-2.24,1.46,0,2.37,5.88,2.16,5.88,5.9C69.15,19.36,67.85,21.07,64.86,21.07Z"></path></g><g fill="#a7a8aa"><path class="b" d="M80,16.46a16.24,16.24,0,0,1-1.88-.1v4.38H76.52V8.65h3.92a3.78,3.78,0,0,1,4.1,3.82C84.54,14.83,82.24,16.46,80,16.46ZM79.9,9.93H78.1v5.12s.72.08,1.28.08c1.78,0,3.36-.6,3.36-2.66C82.74,10.77,81.8,9.93,79.9,9.93Z"></path><path class="b" d="M91.74,13.12l-1.26.36c0-.14-.06-.74-.64-.74-.88,0-1.6,1.46-1.92,2.24v5.76H86.38V14.15a14.56,14.56,0,0,0-.08-1.76l-.08-.72h1.42l.26,1.44A2.48,2.48,0,0,1,90,11.45a1.53,1.53,0,0,1,1.74,1.32c0,.14,0,.22,0,.24Z"></path><path class="b" d="M94,16v.3c0,2.16,1.06,3.34,3,3.34a4.9,4.9,0,0,0,2.84-.94l.12,1.4A6.1,6.1,0,0,1,96.5,21c-2.28,0-4.08-1.76-4.08-4.74s1.7-4.8,4-4.8c2.78,0,3.64,1.44,3.64,4.58Zm4.68-1.2c-.06-1.76-.84-2.08-2.34-2.08a2.91,2.91,0,0,0-1.2.28,2.85,2.85,0,0,0-1,1.8Z"></path><path class="b" d="M105.28,18.4l2.36-6.72h1.52l-3.38,9.06H104.6L101,11.67h1.64Z"></path><path class="b" d="M111.54,10.09a1.11,1.11,0,0,1-1-1.12,1,1,0,1,1,2,0A1.15,1.15,0,0,1,111.54,10.09Zm-.76,1.58h1.54v9.06h-1.54Z"></path><path class="b" d="M116,16v.3c0,2.16,1.06,3.34,3,3.34a4.9,4.9,0,0,0,2.84-.94l.12,1.4a6.1,6.1,0,0,1-3.44.86c-2.28,0-4.08-1.76-4.08-4.74s1.7-4.8,4-4.8c2.78,0,3.64,1.44,3.64,4.58Zm4.68-1.2c-.06-1.76-.84-2.08-2.34-2.08a2.91,2.91,0,0,0-1.2.28,2.85,2.85,0,0,0-1,1.8Z"></path><path class="b" d="M132.24,18.34l2.18-6.66h1.52l-3.14,9.06h-1.16l-2.14-5.68-2,5.68h-1.16L123,11.67h1.64l2.44,6.66,2.44-7.06Z"></path></g></svg></span></a><div id="gh-nav-cnt" class="u-flex-center-ver">
<nav aria-label="links" class="gh-nav gh-nav-links gh-nav-h">
<ul class="gh-nav-list u-list-reset paddingTopHalf">
<li class="gh-nav-item gh-move-to-spine">
<a class="anchor button-link-primary gh-nav-action " href="/freelookup/form/author.uri?zone=TopNavBar&amp;origin=recordpage" id="gh-Authorsearch">
<span class="anchor-text">Author search</span>
</a>
</li>
<li class="gh-nav-item gh-move-to-spine">
<a class="anchor button-link-primary gh-nav-action  " href="/sources?zone=TopNavBar&amp;origin=recordpage" id="gh-Sources">
<span class="anchor-text">Sources</span>
</a>
</li>
</ul>
</nav>   
<nav aria-label="utilities" class="gh-nav gh-nav-utilities gh-nav-h paddingTopHalf">
<ul class="gh-nav-list u-list-reset">
<li class="gh-nav-item gh-has-dd">
<div class="gh-ppvr btn-group" id="helpMenu">
<div class="gh-ppvr-trigger">
<button class="button-link button-medium button-link-primary gh-nav-action gh-icon-btn" id="aa-globalheader-Help" aria-expanded="false" data-toggle="dropdown" type="button">
<svg focusable="false" role="img" viewBox="0 0 114 128" width="21.375" height="24" class="gh-icon icon-help" aria-hidden="true">
<path d="m57 8c-14.7 0-28.5 5.72-38.9 16.1-10.38 10.4-16.1 24.22-16.1 38.9 0 30.32 24.68 55 55 55 14.68 0 28.5-5.72 38.88-16.1 10.4-10.4 16.12-24.2 16.12-38.9 0-30.32-24.68-55-55-55zm0 1e1c24.82 0 45 20.18 45 45 0 12.02-4.68 23.32-13.18 31.82s-19.8 13.18-31.82 13.18c-24.82 0-45-20.18-45-45 0-12.02 4.68-23.32 13.18-31.82s19.8-13.18 31.82-13.18zm-0.14 14c-11.55 0.26-16.86 8.43-16.86 18v2h1e1v-2c0-4.22 2.22-9.66 8-9.24 5.5 0.4 6.32 5.14 5.78 8.14-1.1 6.16-11.78 9.5-11.78 20.5v6.6h1e1v-5.56c0-8.16 11.22-11.52 12-21.7 0.74-9.86-5.56-16.52-16-16.74-0.39-0.01-0.76-0.01-1.14 0zm-4.86 5e1v1e1h1e1v-1e1h-1e1z"></path>
<rect x="0" y="0" width="100%" height="100%" stroke="none" opacity="0">
<title>Help</title>
</rect>
</svg>
<span class="sr-only">Help</span>
</button>
</div>
<div id="gh-ppvr-cnt-helpSubMenu" class="gh-ppvr-cnt gh-ppvr-right dropdown-menu">
<div class="gh-ppvr-arrow-cnt">
<div class="gh-ppvr-arrow gh-arrow-right">
<div class="u-position-relative">
<div class="gh-ppvr-arrow-fill"></div>
</div>
</div>
</div>
<div class="gh-ppvr-cnt-inner">
<h2 class="dropdownHeading sr-only">Help Links</h2>
<ul class="gh-dd-nav">
<li class="gh-nav-item">
<a class="anchor button-link-primary gh-nav-action dropdown-item" href="#" target="globalHelp" onClick="return openGlobalhelp('/standard/help.uri?zone=TopNavBar&amp;origin=recordpage&amp;topic=14190','globalHelp',760,570,0)">
<span class="anchor-text">Help</span>
<span class="ico-external-link-after noMargin" aria-hidden="true"></span>
<span class="sr-only"></span>
</a>
</li>
<li class="gh-nav-item">
<a class="anchor button-link-primary gh-nav-action dropdown-item" href="#" id="gh-Tutorials" target="globalHelp" onClick="return openGlobalhelp('/standard/tutorial.uri?zone=TopNavBar&amp;origin=recordpage&amp;module=revdoc', 'globalHelp',760,570,0)" title="Tutorials">
<span class="anchor-text">Tutorials</span>
<span class="ico-external-link-after noMargin" aria-hidden="true"></span>
<span class="sr-only"></span>
</a>
</li>
<li class="gh-nav-item">
<a class="anchor button-link-primary gh-nav-action dropdown-item" href="#" id="gh-Contactus" target="_blank" onClick="return openGlobalhelp('/standard/contactForm.uri?zone=TopNavBar&amp;origin=recordpage', 'globalHelp',760,570,0)" title="Contact us">
<span class="anchor-text">Contact us</span>
<span class="ico-external-link-after noMargin" aria-hidden="true"></span>
<span class="sr-only"></span>
</a>
</li>
</ul>
</div>
</div>
</div>
</li>
<li class="gh-nav-item gh-institution-item gh-has-dd gh-move-to-spine">
<div class="gh-ppvr u-clr-grey7 text-s btn-group" id="gh-institution-dd">
<div class="gh-ppvr-trigger">
<button class="button-link button-medium button-link-primary gh-nav-action gh-icon-btn" id="aa-globalheader-Institutions" aria-expanded="false" data-toggle="dropdown" type="button"  >
<svg focusable="false" role="img" viewBox="0 0 106 128" width="19.875" height="24" class="gh-icon icon-institution" aria-hidden="true">
<path d="m84 98h1e1v1e1h-82v-1e1h1e1v-46h14v46h1e1v-46h14v46h1e1v-46h14v46zm-72-61.14l41-20.84 41 20.84v5.14h-82v-5.14zm92 15.14v-21.26l-51-25.94-51 25.94v21.26h1e1v36h-1e1v3e1h102v-3e1h-1e1v-36h1e1z"></path>
<rect x="0" y="0" width="100%" height="100%" stroke="none" opacity="0">
<title>Institutions</title>
</rect>
</svg>
<span class="button-link-text">
<span class="sr-only">Institutions</span>
</span>
</button>
</div>
<div id="gh-institution" class="gh-ppvr-cnt gh-ppvr-right dropdown-menu">
<div class="gh-ppvr-arrow-cnt">
<div class="gh-ppvr-arrow gh-arrow-right">
<div class="u-position-relative">
<div class="gh-ppvr-arrow-fill"></div>
</div>
</div>
</div>
<div class="gh-ppvr-cnt-inner">
<div id="gh-institution-body" class="gh-inst-cnt">
<h2 class="gh-inst-lbl u-margin-bottom-s dropdownHeading">Get access via your institution</h2>
<p>Make sure you get full access once you've left your institution.</p>
<p class="u-margin-top-s">Access anytime, anywhere?</p>
<p class="u-margin-top-s">
<a href="/checkaccess.uri?zone=TopNavBar&amp;origin=recordpage" class="btn btn-primary btn-sm gh-flex-center-text guideLinr-element-highlight">
<span class="anchor-text">Check access</span>
</a>
</p>
</div>
</div>
</div>
</div>
</li>
</ul>
</nav>
</div>
<div class="gh-profile-container">
<a href="/signin.uri?origin=recordpage&zone=TopNavBar" id="signin_registerlink" class="btn btn-secondary u-margin-left-s gh-flex-center-text altColor" target="_top">
<span class="btn-text">Create account</span>
<span class="sr-only">Takes you to the ID+ sign in page</span>
</a>
<a href="/signin.uri?origin=recordpage&zone=TopNavBar" id="signin_link_move" class="btn btn-primary gh-flex-center-text" target="_top">
<span class="btn-text">Sign in</span>
</a>
</div>
</div>
<div id="loginPromptWrap" class="hidden">
<div class="dropdownGroup registerSigninPrompt">
<div class="dropdown-menu loginPromptMenu">
<div class="dropdownHeader"></div>
<div class="dropdownContent">
<div class="dropdownContentWrap">
<h4 class="row" tabindex="-1">
Sign in or create account
<a href="/standard/help.uri?topic=21591" class="btn-link" target="globalHelp" 
onclick="return openGlobalhelp('/standard/help.uri?topic=21591','globalHelp',760,570,0)" title="Learn more about signing in or creating an account with Scopus (opens in a new window)">
<span class="sr-only">Learn more about signing in or creating an account with Scopus (opens in a new window)</span>
<span class="ico-help-icon fontMedium" aria-hidden="true"></span>
</a>
</h4>
<p>To use this feature you must be a registered user of Scopus.</p>
<button class="btn btn-primary btn-sm pull-right dropdownActionButton" type="button" title="Sign in or create an account with Scopus">
<span class="btnText">Sign in or create account</span>
</button>
<button class="btn btn-link closeDropDown" data-toggle="dropdown" type="button">
<span class="ico-delete icon" aria-hidden="true"></span>
<span class="sr-only">Close popup</span>
</button>
</div>
</div>
<span class="dropdownFooter"></span>
</div>
</div>
</div>
</header>
<div class="row">
<div class="col-md-12">
<sc-page-title pageTitle="Document details - An efficient highly parallelized ReRAM-based architecture for motion estimation of HEVC" data-sr-only=""></sc-page-title>
<script defer src="https://components.scopus.com/www/scopus-page-title/scopus-page-title.js"></script>
</div>
</div>							
<div class="row">
<div id="contentWrapper" class="col-md-12">
<div id="platformAlertArea" aria-live="polite" aria-atomic="true"></div>
<div id="alertArea" aria-live="polite" aria-atomic="true" >
<noscript>
<div class="alert alert-danger" role="alert">
<div class="jsMsgImg">
<img src="/static/images/warning_small.gif" />
</div>
<div class="jsMsgTxt">
<span class="Bold">We have detected that your browser has JavaScript disabled.</span>
<br>
<span>Scopus needs JavaScript to function correctly. Please turn on JavaScript to continue.</span>
</div>
</div>
</noscript>
<div class="alert alert-danger ie9InfoMsgCnt displayNone" role="alert">
<a href="#" class="close" data-dismiss="alert" aria-label="close">×</a>
<p>Scopus will soon cease the support of IE 9 and users are recommended to upgrade to the latest Internet Explorer, Firefox, or Chrome.</p>
</div>
</div>
<div id="container"> 
<a name="skip1" id=skip1 href="#"></a>
<div class="row wrapper whiteBg">
<div id="prevProfilelayout" class="row">	
<div id="profileleftside"  class="col-md-9">	
<div id="docDetailPage">
<section id="topResultListLinks1">
<nav class="navigationLink" id="refNavLinksId">
<ul class="list-inline">
<li class="recordPageCount"> 
<strong>1  of 1</strong>
</li>
</ul>
</nav>
</section>
<section id="quickLinks">
<ul class="list-inline">
<li id="exportFeature">
<a class="disabled" title='Export this document: Subscription required'>
<span class="ico-log-in" aria-hidden="true"></span>
<span class="anchorText">Export</span>
</a>
</li>
<li id="downloadFeature">
<a class="disabled" title='Download the abstract or full text of the selected document(s), depending on availability: Subscription required'>
<span class="ico-download" aria-hidden="true"></span>
<span class="anchorText">Download</span>
</a>
</li>				
<li id="moreMenu">
<menu class="dropdownGroup">
<a class="disabled" data-type="moreLink" title="Add to My List, Print this document, Email or Create bibliography">
<span class="anchorText">More...</span>
<span class="ico-navigate-right" aria-hidden="true"></span>
</a>
</menu>
</li>
</ul>
</section>						
</div>	   
<div id="profileleftinside">
<section xmlns:htm="http://www.w3.org/1999/xhtml" xmlns:localzn="xalan://com.elsevier.scopus.biz.util.LocalizationHelper" id="articleTitleInfo" class="list-group marginTop1 noMarginLeft paddingTop1">
<span id="guestAccessSourceTitle" class="list-group-item">Journal of Systems Architecture</span><span id="journalInfo" class="list-group-item">Volume 117, August 2021, Article number 102123</span><span class="list-group-item"></span>
</section><input id="sourceId" type="hidden" value="12398"><input id="sourceType" type="hidden" value="j"><input id="citationType" type="hidden" value="ar"><input id="currentRecordPageEID" type="hidden" value="2-s2.0-85104055853"><input id="currentRecordPageSCP" type="hidden" value="85104055853"><input id="publicationYear" type="hidden" value="2021">
<div>
<h2 xmlns:localzn="xalan://com.elsevier.scopus.biz.util.LocalizationHelper" class="h3">An efficient highly parallelized ReRAM-based architecture for motion estimation of HEVC<span>(Article)</span>
</h2>
<section id="authorlist">
<ul class="list-inline">
<li>
<span class="previewTxt" title="Show author details: Subscription required">Zhang, Y.</span><span class="guestView">,</span>
</li>
<li>
<span class="previewTxt" title="Show author details: Subscription required">Liu, B.</span><span class="guestView">,</span>
</li>
<li>
<span class="previewTxt" title="Show author details: Subscription required">Jia, Z.</span><span class="disabled" title="Email this author: Subscription required"></span><span class="guestView">,</span>
</li>
<li>
<span class="previewTxt" title="Show author details: Subscription required">Chen, R.</span><span class="disabled" title="Email this author: Subscription required"></span><span class="guestView">,</span>
</li>
<li>
<span class="previewTxt" title="Show author details: Subscription required">Shen, Z.</span>
</li>
<li>
<a name="corrAuthorTitle"><span class="ico-person" aria-hidden="true"></span><span class="sr-only">View Correspondence (jump link)</span></a>
</li>
</ul>
</section>
<section id="viewHideAuthorList-dummy" class="previewTxt">
<a class="viewAddIcon" href="#" aria-expanded="false" title="View additional authors"><span class="anchorText">View additional authors</span>&nbsp;
<span class="ico-navigate-down" aria-hidden="true"></span></a>
</section>
<a id="saveAllToAuthorList" class="secondaryLink displayInlineBlock paddingBottom1" href="#" title="Save all to author list"><span class="ico-create-bibliography" aria-hidden="true"></span>&nbsp;
<span class="anchorText">Save all to author list</span></a>
<section id="AuthorRetriving" class="hidden">
<div>
<img src="/static/images/loading_3.gif"></div>
<div>Retrieving additional authors...</div>
</section>
<section id="affiliationlist">
<ul class="list-unstyled">
<li id="superscript_a">
<sup class="guestView">a</sup>School of Computer Science and Technology, Shandong University, China</li>
<li id="superscript_b">
<sup class="guestView">b</sup>College of Intelligence and Computing, Shenzhen Research Institute of Tianjin University, Tianjin University, China</li>
</ul>
</section>
<section id="viewHideAffilList-dummy">
<a id="viewMoreAffilite" href="#" aria-expanded="false" title="View additional affiliations" class="disabled"><span class="anchorText">View additional affiliations</span>&nbsp;
<span class="ico-navigate-down" aria-hidden="true"></span></a>
</section>
<section id="AffRetriving" class="hidden">
<div>
<img src="/static/images/loading_3.gif"></div>
<div>Retrieving additional affiliations...</div>
</section>
<a name="abstract"></a>
<section id="abstractSection" class="row">
<h3 class="h4">Abstract<span class="pull-right hidden"><a title="View reference: Subscription required" class="ico-navigate-down"><span class="anchorText">View references</span></a></span>
</h3>
<p>Motion estimation (ME) is a high efficiency video coding (HEVC) process for determining motion vectors that describe the blocks transformation direction from one adjacent frame to a future frame in a video sequence. ME is a memory and computation consuming process which accounts for more than 50% of the total running time of HEVC. To conquer the memory and computation challenges, this paper presents ReME, a highly paralleled processing-in-memory (PIM) architecture for the ME process based on resistive random access memory (ReRAM). In ReME, the space of ReRAM is mainly separated into storage engine and ME processing engine. The storage engine is used as conventional memory to store video frames and intermediate data, while the computation operations of ME are performed in ME processing engines. Each ME processing engine in ReME consists of Sum of Absolute Differences (SAD) modules, interpolation modules, and Sum of Absolute Transformed Difference (SATD) modules that transfer ME functions into ReRAM-based logic analog computation units. ReME further cooperates these basic computation units to perform ME processes in a highly parallel manner. Simulation results show that the proposed ReME accelerator significantly outperforms other implementations with time consuming and energy saving. © 2021 Elsevier B.V.</p>
</section>
<div id="reaxysSubstanceBlk"></div>
<div id="reaxysBlk"></div>
<section id="authorKeywords">
<h3 class="h4">Author keywords</h3>
<span class="badges">High efficiency video coding (HEVC)</span><span class="badges">Motion estimation</span><span class="badges">Processing-in-memory (PIM)</span><span class="badges">Resistive random access memory (ReRAM)</span>
</section>
<section id="indexedKeywords">
<h3 class="h4">Indexed keywords</h3>
<table class="table">
<tr>
<th scope="row">Engineering controlled terms:</th><td><span class="badges">Analog computers</span><span class="badges">Computation theory</span><span class="badges">Energy conservation</span><span class="badges">Engines</span><span class="badges">Image coding</span><span class="badges">Memory architecture</span><span class="badges">RRAM</span><span class="badges">Video signal processing</span><span class="badges">Virtual storage</span></td>
</tr>
<tr>
<th scope="row">Engineering uncontrolled terms</th><td><span class="badges">Analog computation</span><span class="badges">Conventional memory</span><span class="badges">High-efficiency video coding</span><span class="badges">Highly parallels</span><span class="badges">Processing engine</span><span class="badges">Processing in memory</span><span class="badges">Resistive Random Access Memory (ReRAM)</span><span class="badges">Sum of absolute differences</span></td>
</tr>
<tr>
<th scope="row">Engineering main heading:</th><td><span class="badges">Motion estimation</span></td>
</tr>
<tr>
<td colspan="2"></td>
</tr>
</table>
</section>
<section id="fundingDetails">
<h3 class="h4">Funding details</h3>
<table class="table">
<thead>
<tr>
<th scope="col">
Funding sponsor</th><th scope="col">
Funding number</th><th scope="col">
Acronym</th>
</tr>
</thead>
<tbody>
<tr class="lightGreyBorderBottom">
<td>Science and Technology Foundation of Shenzhen City</td><td>JCYJ20170816093943197</td><td></td>
</tr>
<tr class="lightGreyBorderBottom">
<td>National Natural Science Foundation of China</td><td>61902218</td><td>NSFC</td>
</tr>
</tbody>
</table>
<ul class="list-unstyled">
<li>
<h4 id="fundingTextHead1" class="fundingTextHead">1</h4>
<p id="fundingText">The work described in this paper is partially supported by Shenzhen Science and Technology Foundation ( JCYJ20170816093943197 ), and the grants from the National Science Foundation for Young Scientists of China (Grant No. 61902218 ).</p>
</li>
</ul>
</section>
<section id="referenceInfo">
<div class="row">
<div class="col-md-5">
<ul class="list-unstyled" id="citationInfo">
<li>
<strong>ISSN: </strong>13837621</li>
<li>
<strong>CODEN: </strong>JSARF</li>
<li>
<strong>Source Type: </strong>Journal</li>
<li>
<strong>Original language: </strong>English</li>
</ul>
</div>
<div class="col-md-7">
<ul class="list-unstyled" id="documentInfo">
<li>
<strong>DOI: </strong><span id="recordDOI">10.1016/j.sysarc.2021.102123</span>
</li>
<li>
<strong>Document Type: </strong>Article</li>
<li>
<strong>Publisher: </strong>Elsevier B.V.</li>
</ul>
</div>
</div>
</section>
</div>
<div id="SC_BA1P" class="sgfNoTitleBar svDoNotLink">
</div>
<div id="SC_BA1" class="svDoNotLink">	       
</div>
<a name="ref"></a>
<hr>
<div id="refDocs"></div>
<p class="corrAuthSect">
<a name="corrAuthorFooter"><span class="ico-person">&nbsp;</span></a> Jia, Z.; School of Computer Science and Technology, Shandong University, China; <BR>
&copy; Copyright 2021 Elsevier B.V., All rights reserved.</p>
</div>
</div>
<div id="profilerightside"  class="col-md-3">
<div id="profilerightinside">
<div id='chapterList'  class ="displayNone"></div>
<script type="text/javascript">
var chapterListEnabled = true;
</script>
<div>
<div id="recordPageBoxes">
<div class="panel panel-default boxPanel">
<div class="panel-heading">
<h3 class="panel-title">
Cited by 0 documents 
</h3>
</div>
</div> 
<div class="panel-footer">
<div class="footerSection">
<span class="findMoreHeader ">Inform me when this document is cited in Scopus:</BR></span>
<div class="citeinfo flexDisplay verticalAlign">
<a class="btn btn-secondary-altBg btn-xsm altColor saveAsAlert disabled" title='Set citation alert'>
<span class="anchorText">Set citation alert</span>
<span class="ico-navigate-right" aria-hidden="true"></span>
</a>
<a class="btn btn-secondary-altBg btn-xsm altColor icon setFeed disabled" title='Set citation feed'>
<span class="anchorText">Set citation feed</span>
<span class="ico-navigate-right" aria-hidden="true"></span>
</a>
</div>
</div>      
</div>
			
</div>
</div>
</div>
<script language="javascript" type="text/javascript">
var mltContinueURL;
if (mltContinueURL != null) {
yes = confirm(mltContinueMsg);
if (yes) {
window.location = mltContinueURL;
}
}
</script>
<div style="display:none" id="mlttable">
<div id="recordPageBoxes">
<div id="relatedDocBox" class="panel panel-default boxPanel">
<div class="panel-heading">
<h3 class="panel-title">Related documents</h3>
</div>
<div class="panel-body">
<div id="relatedDoc"></div>
</div>
<div class="panel-footer">
<div class="inforeldocbox footerSection">
<span class="findMoreHeader">Find more related documents in Scopus based on:</span>
<div class="footerSctnLinks flexDisplay verticalAlign">
<div class="pull-left">
<div class="">
<a class="disabled" title='Authors: Subscription required'>
<span class="anchorText">Authors</span>
<span class="ico-navigate-right" aria-hidden="true"></span>
</a>
</div>
</div>
<div class="pull-left keywordsMargin">
<div>
<a class="disabled" title='Keywords: Subscription required'>
<span class="anchorText">Keywords</span>
<span class="ico-navigate-right" aria-hidden="true"></span>
</a>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
<script language="javascript" type="text/javascript">
document.getElementById('mlttable').style.display = "block";
</script>
<section id="articleDataSet"></section>
</div>
</div>
<div class="sciTopicsVal displayNone">{"topic":{"name":"Particle Accelerators; RRAM; TOPS","id":1016560,"uri":"Topic/1016560","prominencePercentile":99.622604,"prominencePercentileString":"99.623","overallScholarlyOutput":0},"dig":"465dab48786fa6b063e7062d66e1669261a66747c7b5c2a0cf56da8f58d28cf9"}</div>
<div class="sciTopics">
<section id="topicSection" class="row">
<h3 class="topicTitle h4 displayInline">SciVal Topic Prominence</h3>
<span class="dropdown dropup">
<button type="button" title="Learn about these Topics" class="btn btn-link dropdown-toggle btn-sm" id="topicsBtn" data-toggle="dropdown"
aria-haspopup="true" aria-expanded="false">
<span class="ico-information triangleColor" aria-hidden="true"></span>
</button>
<div class="dropdown-menu dropdown-menu-right dropdown-menu-medium" role="menu" aria-labelledby="downloadHelp" id="topicsHelp">
<span class="dropdownHeader"></span>
<span class="dropdownContent">
<span class="dropdownContentWrap">
<span id="dropdownTitle" class="row" tabindex="-1"></span>
<h4 id="meModalTitle"><strong class="font-weight-bold">New: </strong>SciVal Topic Prominence</h4>
<p>Topics are unique areas of research, created using all Scopus publications from 1996 onwards. </p>
<p class="noMargin">
Use this section to learn about the Topic, find key authors to follow, and view related documents.
<a class="btn btn-sm secondaryLink externalLink marginTopHalf noBorder" href="/standard/help.uri?topic=27947"
alt="" target="globalHelp" title="Learn more about these Topics (opens in a new window)" id="topicDetail"> 
<span class="anchorText paddingTopHalf">Learn more about these Topics</span>
<span class="ico-external-link-after" aria-hidden="true"></span>
</a>
</p>
<button class="btn btn-link closeDropDown" data-toggle="dropdown" id="closeDownloadHelpMenu" type="button">
<span class="ico-delete icon" aria-hidden="true"></span>
<span class="sr-only"> Close window </span>
</button>
</span>
</span>
<span class="dropdownFooter"></span>
</div>
</span>
<div class="topicContent marginBottom1 marginTop1">
Topic: 
<a href="#topicDetailModal" data-toggle="modal" data-target="#topicDetailModal" title="Open Topic details" id="topicLink" class="paddingLeftHalf secondaryLink"></a>
</div>
<div class="topicsProgress flexAlignCenter flexDisplay">
<div>
Prominence percentile:
<span class="paddingRightHalf paddingLeft1 percentText"></span>
</div>
<div class="progress quarterWidth noMarginBottom" style="width: 4.9rem">
<div class="progress-bar">&nbsp;</div>
</div>
<div class="dropdown dropup">
<button type="button" title="Learn about prominence percentile"
class="btn btn-link dropdown-toggle btn-sm marginLeftHalf" id="svTopicInfo"
data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
<span class="ico-information triangleColor" aria-hidden="true"></span>
</button>
<div class="dropdown-menu dropup-menu-right dropdown-menu-medium marginLeft1" role="menu" aria-labelledby="svTopicInfo">
<span class="dropdownHeader"></span> <span class="dropdownContent">
<span class="dropdownContentWrap greyText">
<span class="">Prominence is an indicator that shows the current momentum of a Topic. It is calculated by weighing 3 metrics for papers clustered in a Topic: Citation count, Scopus views and Average CiteScore.</span>
<button class="btn btn-link closeDropDown" data-toggle="dropdown" id="closeBtnTopicId" type="button">
<span class="ico-delete icon" aria-hidden="true"></span>
</button>
</span>
</span> <span class="dropdownFooter"></span>
</div>
</div>
</div>
</section></div>
<script type="text/javascript" src="/gzip_N1377736354/bundles/d3.js" async="async" ></script>
<div class="modal popupModal" id="topicDetailModal" data-backdrop="true" aria-labelledby="" aria-describedby="" tabindex="-1"
role="dialog">
<div class="modal-dialog modal-xl smallText" role="document">
<div class="modal-content">
<div class="modal-header"><h3><span class="topicName">Topic: </span></h3>
<div><span>Year range: </span><span class="yrRange"></span></div>
<button type="button" class="close" data-dismiss="modal" aria-label="">
<span class="ico-delete icon" aria-hidden="true"></span>
</button>
</div>
<div class="modal-body paddingTop1 noPaddingBottom">
<div class="row">
<div class="col-md-7 pull-left">
<section id="documents">
<header class="paddingBottomHalf fontMedium">Representative documents</header>
<ul class="articleList list-unstyled"></ul>
</section>
</div>
<div class="col-md-5 pull-right">
<aside id="authors" class="col-md-10">
<table class="table">
<thead>
<th class="noBorder noPaddingLeft noPaddingTop fontMedium paddingBottomHalf">Top authors</th>
<th class="noBorder textRight noPaddingLeft noPaddingTop lightGreyText fontSmall paddingBottomHalf">Scholarly Output</th>
</thead>
<tbody>
</tbody>
</table>
</aside>
<aside id="visual" class="col-md-12">
<div id="tabs" class="clearfix marginTop2 marginBottom1Half marginRight1">
<header class="pull-left fontMedium">Keyphrase analysis</header>
<ul class="noBorder pull-right nav buttonTab">
<li data-tabname="wrdCldCtner" class="pull-left active" default>
<button href="#wrdCldCtner" class="btn btn-sm paddingRightHalf paddingLeftHalf btn-secondary btn-link secondary-button-border-color noBorderRight" data-toggle="tab" id="chartTab">
<span class="ico-upwards-line-chart paddingRightQuarter"></span><span class="btnText noBorder">Chart</span>
</button>
</li>
<li data-tabname="keyword" class="pull-left">
<button href="#keyword" class="btn btn-sm paddingRightHalf paddingLeftHalf btn-secondary btn-link secondary-button-border-color noBorderLeft" data-toggle="tab" id="tableTab">
<span class="ico-tables paddingRightQuarter"></span><span class="btnText noBorder">Table</span>
</button>
</li>
</ul>
</div>
<div class="tab-content">
<div id="wrdCldCtner" class="tab-pane fade in active">
<div id="chart"></div>
<div class="wordCloudLegend textCenter">
<span class="relevance">
<span class="normal">A</span>
<span class="large">A</span>
<span class="larger">A</span>
<span>relevance of keyphrase</span>
<span class="verticalpipe">|</span>
</span>
<span class="nowrap">
<span>declining</span>
<span class="large declining">A</span>
<span class="large flat">A</span>
<span class="large growing">A</span>
<span>Growth</span>
</span>
</div>
</div>
<div id="keyword" class="tab-pane">
<table id="keyword-table" class="table table-fixed tableRowBorder"></table>
</div>
</div>
</div>
</aside>
</div>
</div>
<div class="row pull-right paddingRight1 paddingBottom1 marginTop2">
<a class="btn btn-primary btn-sm" data-link='https://scival.com/viewtopic?id=' title="Analyze this topic in SciVal (opens in new tab)"
id="topicsLink" target="_blank"> <span class="btnText paddingRightHalf noBorder">Analyze in SciVal</span><span class="ico-navigate-right whiteText"
aria-hidden="true"></span>
</a>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
<div class="clear"></div>
</div>
</div>
<div class="row">
<div class="col-md-12 footer">
<footer>
<a id="skiptofooter"></a>
<div class="row footerLinks">
<div class="col-md-3">
<h3>About Scopus</h3>
<ul class="list-unstyled">
<li>
<a target="_blank" href="https://www.elsevier.com/online-tools/scopus?dgcid=RN_AGCM_Sourced_300005030" title='Learn more about Scopus (opens in a new window)' class="secondaryLink">
<span class="anchorText">What is Scopus</span>
</a>
</li>
<li>
<a target="_blank" href="https://www.elsevier.com/online-tools/scopus/content-overview/?dgcid=RN_AGCM_Sourced_300005030" title='Learn more about Scopus&#039; content coverage (opens in a new window)' class="secondaryLink">
<span class="anchorText">Content coverage</span>
</a>
</li>
<li>
<a target="_blank" href="https://blog.scopus.com/" title='Read the Scopus Blog (opens in a new window)' class="secondaryLink">
<span class="anchorText">Scopus blog</span>
</a>
</li>
<li>
<a target="_blank" href="https://dev.elsevier.com/" title='Learn more about Scopus API&#039;s (opens in a new window)' class="secondaryLink">
<span class="anchorText">Scopus API</span>
</a>
</li>
<li>
<a target="_blank" href="https://www.elsevier.com/about/our-business/policies/privacy-principles?dgcid=RN_AGCM_Sourced_300005030" title='View privacy matters page (opens in a new window)' class="secondaryLink">
<span class="anchorText">Privacy matters</span>
</a>
</li>
</ul>
</div>
<div class="col-md-3">
<h3>Language</h3>
<ul class="list-unstyled" id="languageSwitchList" data-currentlocale="en_US">
<li lang="ja" >
<a lang="ja" hreflang="ja" href="/personalization/switch/Japanese.uri?origin=recordpage&zone=footer&locale=ja_JP" target="_top" title="日本語版を表示する" class="">
<span class="anchorText">
日本語に切り替える
</span>
</a>
</li>
<li lang="zh-Hans" >
<a lang="zh-Hans" hreflang="zh-Hans" href="/personalization/switch/Chinese.uri?origin=recordpage&zone=footer&locale=zh_CN" target="_top" title="查看简体中文版本" class="">
<span class="anchorText">
切换到简体中文
</span>
</a>
</li>
<li lang="zh-Hant" >
<a lang="zh-Hant" hreflang="zh-Hant" href="/personalization/switch/Chinese.uri?origin=recordpage&zone=footer&locale=zh_TW" target="_top" title="查看繁體中文版本" class="">
<span class="anchorText">
切換到繁體中文
</span>
</a>
</li>
<li lang="ru" >
<a lang="ru" hreflang="ru" href="/personalization/switch/Russian.uri?origin=recordpage&zone=footer&locale=ru_RU" target="_top" title="Просмотр версии на русском языке" class="">
<span class="anchorText">
Русский язык
</span>
</a>
</li>
</ul>
</div>
<div class="col-md-6">
<h3>Customer Service</h3>
<ul class="list-unstyled">
<li>
<a target="_blank" title="View Scopus help files (opens in a new window)" href="/standard/contactUs.uri?pageOrigin=footer">
<span class="anchorText">Help</span>
</a>
</li>
<li>
<a target="_blank" title="Contact us (opens in a new window)" 
href="/standard/contactForm.uri?pageOrigin=footer">
<span class="anchorText">
Contact us
</span>
</a>
</li>
</ul>
</div>
</div>
<div class="row footerCopyright">
<div class="col-md-2">
<a id="elsevierLogo" target="_blank" href="https://www.elsevier.com/?dgcid=RN_AGCM_Sourced_300005030" title="Go to the Elsevier site (opens in a new window)">
<img src="/static/images/logo_Elsevier.svg" alt="Elsevier">
</a>
</div>
<div class="col-md-8">
<ul class="list-unstyled list-inline">
<li>
<a target="_blank" href="https://www.elsevier.com/locate/termsandconditions?dgcid=RN_AGCM_Sourced_300005030" title="View the terms and conditions of Elsevier (opens in a new window)" class="btn-link secondaryLink noPaddingLeft noBorder">
<span class="anchorText">
Terms and conditions
</span>
<span class="ico-external-link-after noMargin" aria-hidden="true"></span>
<span class="sr-only">View the terms and conditions of Elsevier (opens in a new window)</span>
</a>
</li>
<li>
<a target="_blank" href="https://www.elsevier.com/locate/privacypolicy?dgcid=RN_AGCM_Sourced_300005030" title="View the privacy policy of Elsevier (opens in a new window)" class="btn-link secondaryLink noPaddingLeft noBorder">
<span class="anchorText">
Privacy policy
</span>
<span class="ico-external-link-after noMargin" aria-hidden="true"></span>
</a>
</li>
</ul>
<p id="copyRightID" class="noMargin">
Copyright ©  <a target="_blank" href="https://www.elsevier.com?dgcid=RN_AGCM_Sourced_300005030" title="Go to the Elsevier site (opens in a new window)" class="btn-link noPadding secondaryLink noBorder"><span class="anchorText">Elsevier B.V </span><span class="ico-external-link-after noMargin" aria-hidden="true"></span></a>. All rights reserved. Scopus® is a registered trademark of Elsevier B.V.
</p>
<p class="cookiesTextID noMargin">
We use cookies to help provide and enhance our service and tailor content. By continuing, you agree to the <a href="/cookies/policy.uri" title="See cookie details (opens in a new window)" class="btn-link noPadding secondaryLink">use of cookies</a>.
</p>
</div>
<div class="col-md-2">
<a id="relexLogo" href="http://www.relx.com" title="Go to RELX Group Homepage (Opens in a new window)" target="_blank">
<img src="/static/images/reLogo_orange.svg" alt="RELX Group" />
<img src="/static/images/reText_gray.svg" alt="RELX Group" />
</a>
</div>
</div>
</footer>
<script type="text/javascript" src="/gzip_896231183/bundles/ScopusMasterJS.js" ></script>
<script type="text/javascript" src="/gzip_734864571/bundles/RecordPageBottomMaster.js" ></script>
</div>
</div>
</div>
<input type="hidden" value="278641" id="accId" name="accId">
<script type="text/javascript">	
var WAM = new Object;
WAM.sep = "#";
WAM.ud = "GUEST#289839#278641";
WAM.cars = "";
WAM.part = "HzOxMe3b";
WAM.exportAction=function(format, info, count) {return true;};
</script>
<script type="text/javascript"></script>
<div class="modal fade" id="loadingModal" tabindex="-1" role="dialog" aria-busy="true" aria-live="assertive">
<div class="modal-dialog" role="document">
<div class="modal-content">
<div class="modal-body">
Loading...
<span class="loading-wrapper">
<span class="loadingIcon"></span>
</span>			
</div>
</div>
</div>
</div>
</div>
</div>
<script type="text/javascript" src="/gzip_N1421293715/bundles/pendoTop.js" ></script>
<script type="text/javascript">
var pendoData = {};
pendoDataUtil.loadPendoDataObject("13341253",
"2.44.73.40",
"ae:ANON::GUEST:",
"",
"en_US",
"278641",
"Scopus Preview",
"",
"700DB85889C256DC6548C555F9984E93.i-061e2a7b9e7281261");
(function(p, e, n, d, o) {
var v, w, x, y, z;
o = p[d] = p[d] || {};
o._q = [];
v = [ 'initialize', 'identify', 'updateOptions', 'pageLoad' ];
for (w = 0, x = v.length; w < x; ++w)
(function(m) {
o[m] = o[m]
|| function() {
o._q[m === v[0] ? 'unshift' : 'push']([ m ]
.concat([].slice.call(arguments, 0)));
};
})(v[w]);
y = e.createElement(n);
y.async = !0;
y.src = 'https://content.pendo.scopus.com/agent/static/7108b796-60e0-44bd-6a6b-7313c4a99c35/pendo.js';
z = e.getElementsByTagName(n)[0];
z.parentNode.insertBefore(y, z);
})(window, document, 'script', 'pendo');
/* Call this whenever information about your visitors becomes available. 
* Please use Strings, Numbers, or Bools for value types. */
pendo.initialize({
apiKey : '7108b796-60e0-44bd-6a6b-7313c4a99c35',
visitor : {
id : pendoData.userId,
scopusOriginalId: pendoData.scopusOriginalId,
ipAddress : pendoData.ipAddress,
loginStatus : pendoData.loginStatus,
language : pendoData.language,
scopusLanguage : pendoData.scopusLanguage,
SCSessionID : pendoData.SCSessionID,
navigatorLanguage : window.navigator.language
},
account : {
id : pendoData.accountId,
name : pendoData.accountName,
accessType : pendoData.accessType,
consortiumId : pendoData.consortiumId
}
});
</script>
<script type="text/javascript">
window.appData = window.appData || [];
pageData.event = 'pageLoad';
pageData.page.loadTime = performance ? Math.round(performance.now()).toString() : '';
function reportPageLoad() {
// loadEventEnd metric is not available right after 'load' event
// we have to read it on the next tick, otherwise, it's 0
setTimeout(function () {
if (performance && performance.timing) {
var loadDuration = performance.timing.loadEventEnd - performance.timing.navigationStart;
pageData.page.customPerformance1 = loadDuration.toString();
}
appData.push(pageData);
}, 0);
}
function whenReady() {
if (document.readyState !== 'complete' || !window.pageDataUtilReady) return;
pageDataUtil.loadFooterPageDataObject();
siteCatPageUtil.afterPageLoad();
reportPageLoad();
}
document.addEventListener('pageDataUtilReady', whenReady);
window.addEventListener('load', whenReady);
whenReady();
</script>
<script type="text/javascript">window.NREUM||(NREUM={});NREUM.info={"errorBeacon":"bam-cell.nr-data.net","licenseKey":"0268925da8","agent":"","beacon":"bam-cell.nr-data.net","applicationTime":205,"applicationID":"31454162","transactionName":"MwAGYhBYX0oEABUNXQpKN0YQUF9eJgwPEEALCQhTEBZDXAYMEwAdAAwXRg5YSBlNJCQwGw==","queueTime":0}</script><script defer src="https://static.cloudflareinsights.com/beacon.min.js/v64f9daad31f64f81be21cbef6184a5e31634941392597" integrity="sha512-gV/bogrUTVP2N3IzTDKzgP0Js1gg4fbwtYB6ftgLbKQu/V8yH2+lrKCfKHelh4SO3DPzKj4/glTO+tNJGDnb0A==" data-cf-beacon='{"rayId":"6b5530fbfffb59fb","token":"aad0e745b9764a2aabd77c706b20d432","version":"2021.11.0","si":100}' crossorigin="anonymous"></script>
</body>
</html>
`
