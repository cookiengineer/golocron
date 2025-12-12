
import hljs                from "./highlight.mjs";
import language_bash       from "./languages/bash.mjs";
import language_c          from "./languages/c.mjs";
import language_cpp        from "./languages/cpp.mjs";
import language_css        from "./languages/css.mjs";
import language_diff       from "./languages/diff.mjs";
import language_dns        from "./languages/dns.mjs";
import language_go         from "./languages/go.mjs";
import language_http       from "./languages/http.mjs";
import language_ini        from "./languages/ini.mjs";
import language_javascript from "./languages/javascript.mjs";
import language_json       from "./languages/json.mjs";
import language_markdown   from "./languages/markdown.mjs";
import language_plaintext  from "./languages/plaintext.mjs";
import language_powershell from "./languages/powershell.mjs";
import language_rust       from "./languages/rust.mjs";
import language_sql        from "./languages/sql.mjs";
import language_wasm       from "./languages/wasm.mjs";
import language_nasm       from "./languages/x86asm.mjs";
import language_xml        from "./languages/xml.mjs";
import language_yaml       from "./languages/yaml.mjs";

hljs.registerLanguage('bash',       language_bash);
hljs.registerLanguage('c',          language_c);
hljs.registerLanguage('cpp',        language_cpp);
hljs.registerLanguage('css',        language_css);
hljs.registerLanguage('diff',       language_diff);
hljs.registerLanguage('dns',        language_dns);
hljs.registerLanguage('go',         language_go);
hljs.registerLanguage('http',       language_http);
hljs.registerLanguage('ini',        language_ini);
hljs.registerLanguage('javascript', language_javascript);
hljs.registerLanguage('json',       language_json);
hljs.registerLanguage('markdown',   language_markdown);
hljs.registerLanguage('plaintext',  language_plaintext);
hljs.registerLanguage('powershell', language_powershell);
hljs.registerLanguage('rust',       language_rust);
hljs.registerLanguage('sql',        language_sql);
hljs.registerLanguage('wasm',       language_wasm);
hljs.registerLanguage('nasm',       language_nasm);
hljs.registerLanguage('xml',        language_xml);
hljs.registerLanguage('yaml',       language_yaml);

export const Init = () => {

	return new Promise((resolve, reject) => {

		resolve(hljs);

	});

};
