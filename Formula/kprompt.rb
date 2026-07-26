# typed: false
# frozen_string_literal: true

class Kprompt < Formula
  desc "AI Kubernetes CLI: natural language to a reviewable plan, then approve before apply"
  homepage "https://kprompt.ai"
  version "0.5.0"
  license "Apache-2.0"

  on_macos do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.5.0/kprompt_0.5.0_darwin_arm64.tar.gz"
      sha256 "605c61b2a518dd9e3fdf1fcd386f456e384c768b2ed7e7dcd714830c8b3932e0"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.5.0/kprompt_0.5.0_darwin_amd64.tar.gz"
      sha256 "8ee13ea0553185203f7d36382de3b9ad7255b2cb98aca9ac86e3954c8ab7760b"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/kprompt/kprompt/releases/download/v0.5.0/kprompt_0.5.0_linux_arm64.tar.gz"
      sha256 "6622469b281246c125497ba50b22a22c98358e05f8746ccbb8647bb390fdec60"
    end
    on_intel do
      url "https://github.com/kprompt/kprompt/releases/download/v0.5.0/kprompt_0.5.0_linux_amd64.tar.gz"
      sha256 "85fe357026a3e2487ef7d655a39d5d43b08684fd014052cb918704ba29a1023f"
    end
  end

  def install
    bin.install "kprompt"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/kprompt version")
  end
end
