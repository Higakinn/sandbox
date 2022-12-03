alias ll='ls -al'
alias m='make'
source /usr/share/bash-completion/bash_completion
# json 同士の差分を表示するコマンド
function jsondiff() {
	local file1="${1}"
	local file2="${2}"

	diff <(jq --sort-keys . < "${file1}") <(jq --sort-keys . < "${file2}") | colordiff
}

# file同士の和集合を表示する関数
function union() {
	local a="${1}"
	local b="${2}"

	sort $a $b | uniq
}

# file同士の共通集合を表示する関数
function intersect() {
	local a="${1}"
	local b="${2}"

	sort $a $b | uniq -d
}

function set_diff() {
	local a="${1}"
	local b="${2}"

	sort $a $b $b | uniq -u

}

source <(kubectl completion bash)
