~w(rel plugins *.exs)
|> Path.join()
|> Path.wildcard()
|> Enum.map(&Code.eval_file(&1))

use Mix.Releases.Config,
    default_release: :default,
    default_environment: Mix.env()

environment :dev do
  set dev_mode: true
  set include_erts: false
  set cookie: :"]cl1a]{be$!o,,Ann{w58>}N*aRYmvp,}<0i>Mu.|@Y^0y@0nhBY1d`g>:a9~YA,"
end

environment :prod do
  set include_erts: true
  set include_src: false
  set cookie: :"CwOYa2w(KB3;[L=^gFWqgS`A8;R_D=WpzbkW6E@KPM<mS;IO*_=(sz2_GO&%N]5L"
  set vm_args: "rel/vm.args"
end
