# typed: false
# frozen_string_literal: true

class Kprompt < Formula
  desc "AI Kubernetes CLI: natural language to a reviewable plan, then approve before apply"
  homepage "https://kprompt.ai"
  version "0.9.0"
  license "Apache-2.0"

  on_macos do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.9.0/kprompt_0.9.0_darwin_arm64.tar.gz"
      sha256 "15efdd5fc0d2a67192c72ac9039e224a9a853f2202e4abba05b38365cc420db2"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.9.0/kprompt_0.9.0_darwin_amd64.tar.gz"
      sha256 "95389a3899a89fc24d393aec8fbd0f8d5bf3ab6d39f607c77653bfd68bd41c90"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.9.0/kprompt_0.9.0_linux_arm64.tar.gz"
      sha256 "16240f5e4ce2f3caa73341af0ace607660e851ce9193443d0f517de6ef97af05"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.9.0/kprompt_0.9.0_linux_amd64.tar.gz"
      sha256 "2e105ad8f88777c71ac3a01c16481d16ecd32e3725b2d13ba33473b7efc4a445"
    end
  end

  def install
    bin.install "kprompt"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/kprompt version")
  end
end
