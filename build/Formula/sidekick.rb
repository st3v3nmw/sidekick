class Sidekick < Formula
  desc "A CLI tool that understands your natural language commands and executes them for you"
  homepage "https://github.com/stephen/sidekick"
  version "0.1.0"

  depends_on "go" => :build

  def install
    system "go", "get", "github.com/stephen/sidekick", "./cmd/sidekick"
    system "go", "build", "-o", "sidekick"

    bin.install "sidekick"
  end
end
