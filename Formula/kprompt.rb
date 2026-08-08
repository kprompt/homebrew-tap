# typed: false
# frozen_string_literal: true

class Kprompt < Formula
  desc "AI Kubernetes CLI: natural language to a reviewable plan, then approve before apply"
  homepage "https://kprompt.ai"
  version "0.10.0"
  license "Apache-2.0"

  on_macos do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.10.0/kprompt_0.10.0_darwin_arm64.tar.gz"
      sha256 "5bade31e3f919267a9fd12089bd5ac540ca2716b6684456d4d8039b6418e8e78"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.10.0/kprompt_0.10.0_darwin_amd64.tar.gz"
      sha256 "1d1ffcb6bc10e1e9701ef96e1cd9a481fe85f9a51ff0108cacce8848cc4097c8"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.10.0/kprompt_0.10.0_linux_arm64.tar.gz"
      sha256 "36c1d74f687238545faee787644bbb80c1369d06972ce6facdf58533c4b80b37"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.10.0/kprompt_0.10.0_linux_amd64.tar.gz"
      sha256 "6f942ec173862fe79f825bc2b02bbb7e6e7c0a860485da13b7dc617f8615a78f"
    end
  end

  def install
    bin.install "kprompt"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/kprompt version")
  end
end
