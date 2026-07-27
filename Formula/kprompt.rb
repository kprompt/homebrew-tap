# typed: false
# frozen_string_literal: true

class Kprompt < Formula
  desc "AI Kubernetes CLI: natural language to a reviewable plan, then approve before apply"
  homepage "https://kprompt.ai"
  version "0.6.0"
  license "Apache-2.0"

  on_macos do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.6.0/kprompt_0.6.0_darwin_arm64.tar.gz"
      sha256 "c0ab783c81ecaf8e77817e665d3886b763646b9fe778a7fff907a67595b770ac"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.6.0/kprompt_0.6.0_darwin_amd64.tar.gz"
      sha256 "9ec715d634bc8ed5ffed123b27c172932573b6fe64220b08c0bf1a79dfec844e"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.6.0/kprompt_0.6.0_linux_arm64.tar.gz"
      sha256 "37eb519c1e10e746bf4585acc5fe501e61491b74505f93e02f378549def25c0e"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.6.0/kprompt_0.6.0_linux_amd64.tar.gz"
      sha256 "baadb3b46a2e03e0b85fd8cb6d92947a149b83c63bd094216be66ee42cb57069"
    end
  end

  def install
    bin.install "kprompt"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/kprompt version")
  end
end
