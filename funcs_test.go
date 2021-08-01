package ozon

import (
	"testing"
)

const (
	test1BeforeCut = "  Правильный     ответ  "
	test1AfterCut  = "Правильный ответ "
	test2BeforeCut = "Тестовая        СтРоКа            123           456         Тестовая         строка"
	test2AfterCut  = "Тестовая СтРоКа 123 456 Тестовая строка"
	test3BeforeCut = "Lorem      ipsum      dolor      sit      amet,      consectetur      adipiscing      elit.      Vivamus      eu      consectetur      sem,      a      mattis      ipsum.      Donec      felis      ligula,      rhoncus      at      tempor      vitae,      pretium      eu      tellus.      Donec      accumsan      finibus      facilisis.      Etiam      fringilla      mattis      lobortis.      Aliquam      lobortis      arcu      nisi,      eget      pretium      eros      venenatis      sit      amet.      In      dui      tellus,      fermentum      eget      lorem      non,      aliquet      pharetra      tortor.      Pellentesque      ultrices      sapien      finibus      leo      eleifend,      sit      amet      pellentesque      diam      mattis.\n\nProin      a      velit      et      nisl      tristique      eleifend.      Quisque      vel      elementum      arcu.      Etiam      scelerisque      volutpat      enim      sed      viverra.      Phasellus      eu      massa      in      nunc      auctor      fermentum.      Donec      lobortis,      dolor      non      lobortis      vehicula,      arcu      tellus      vehicula      lectus,      in      luctus      est      sem      eu      eros.      Nulla      in      dictum      est.      Cras      malesuada      eu      purus      mattis      volutpat.      Curabitur      dictum      bibendum      porttitor.      Sed      non      facilisis      ex.      Quisque      dignissim      urna      nec      libero      ullamcorper      finibus.      Sed      eu      gravida      diam.\n\nDonec      faucibus      sodales      tortor      id      volutpat.      Nullam      eu      sem      eu      tellus      rutrum      vehicula.      In      viverra      faucibus      dui      eget      semper.      Curabitur      venenatis,      libero      eu      malesuada      efficitur,      sem      justo      rutrum      ex,      et      sagittis      eros      diam      fermentum      enim.      Ut      id      sagittis      tortor.      Donec      sed      tellus      consectetur,      finibus      mi      vel,      lacinia      metus.      Donec      sit      amet      tellus      sem.      Morbi      sed      vestibulum      ligula,      eget      pharetra      enim.      Vestibulum      porttitor      pellentesque      enim      et      dictum.      Vivamus      non      augue      lectus.      Quisque      lorem      orci,      lacinia      ut      risus      id,      dapibus      pulvinar      enim.      Interdum      et      malesuada      fames      ac      ante      ipsum      primis      in      faucibus.      Phasellus      quis      pellentesque      nisl,      vel      faucibus      arcu.      Pellentesque      lectus      nibh,      feugiat      id      ultricies      non,      sagittis      sit      amet      ex.\n\nSed      a      mauris      maximus,      pellentesque      metus      sit      amet,      tempus      enim.      Integer      suscipit      odio      at      mattis      posuere.      Nulla      vestibulum      elit      nunc.      Pellentesque      habitant      morbi      tristique      senectus      et      netus      et      malesuada      fames      ac      turpis      egestas.      Ut      posuere,      sem      ac      luctus      pharetra,      eros      justo      mattis      metus,      sed      dictum      ante      quam      eget      orci.      Praesent      varius      neque      sem,      ultrices      luctus      justo      vestibulum      et.      Sed      feugiat      quam      mauris.      Class      aptent      taciti      sociosqu      ad      litora      torquent      per      conubia      nostra,      per      inceptos      himenaeos.      Quisque      interdum      orci      sit      amet      ipsum      semper      vehicula.      Vivamus      maximus      rutrum      neque,      non      ultrices      metus      semper      nec.      Nulla      tempus      quam      tristique      ipsum      facilisis,      nec      gravida      diam      mattis.\n\nVivamus      lorem      ex,      efficitur      sed      justo      consectetur,      luctus      dignissim      orci.      Vestibulum      rutrum      pharetra      elit,      in      sollicitudin      tellus      euismod      eget.      Nam      nec      aliquam      metus.      Suspendisse      tristique      ac      massa      ut      fermentum.      Cras      id      mi      eu      sapien      rutrum      ultricies.      Phasellus      id      augue      in      elit      hendrerit      ornare      ac      in      nisl.      In      viverra      leo      vel      elit      posuere,      in      tincidunt      arcu      sodales.      Phasellus      commodo      justo      ac      egestas      pharetra.      Phasellus      at      ante      ut      dui      egestas      faucibus."
	test3AfterCut  = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus eu consectetur sem, a mattis ipsum. Donec felis ligula, rhoncus at tempor vitae, pretium eu tellus. Donec accumsan finibus facilisis. Etiam fringilla mattis lobortis. Aliquam lobortis arcu nisi, eget pretium eros venenatis sit amet. In dui tellus, fermentum eget lorem non, aliquet pharetra tortor. Pellentesque ultrices sapien finibus leo eleifend, sit amet pellentesque diam mattis.\n\nProin a velit et nisl tristique eleifend. Quisque vel elementum arcu. Etiam scelerisque volutpat enim sed viverra. Phasellus eu massa in nunc auctor fermentum. Donec lobortis, dolor non lobortis vehicula, arcu tellus vehicula lectus, in luctus est sem eu eros. Nulla in dictum est. Cras malesuada eu purus mattis volutpat. Curabitur dictum bibendum porttitor. Sed non facilisis ex. Quisque dignissim urna nec libero ullamcorper finibus. Sed eu gravida diam.\n\nDonec faucibus sodales tortor id volutpat. Nullam eu sem eu tellus rutrum vehicula. In viverra faucibus dui eget semper. Curabitur venenatis, libero eu malesuada efficitur, sem justo rutrum ex, et sagittis eros diam fermentum enim. Ut id sagittis tortor. Donec sed tellus consectetur, finibus mi vel, lacinia metus. Donec sit amet tellus sem. Morbi sed vestibulum ligula, eget pharetra enim. Vestibulum porttitor pellentesque enim et dictum. Vivamus non augue lectus. Quisque lorem orci, lacinia ut risus id, dapibus pulvinar enim. Interdum et malesuada fames ac ante ipsum primis in faucibus. Phasellus quis pellentesque nisl, vel faucibus arcu. Pellentesque lectus nibh, feugiat id ultricies non, sagittis sit amet ex.\n\nSed a mauris maximus, pellentesque metus sit amet, tempus enim. Integer suscipit odio at mattis posuere. Nulla vestibulum elit nunc. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut posuere, sem ac luctus pharetra, eros justo mattis metus, sed dictum ante quam eget orci. Praesent varius neque sem, ultrices luctus justo vestibulum et. Sed feugiat quam mauris. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Quisque interdum orci sit amet ipsum semper vehicula. Vivamus maximus rutrum neque, non ultrices metus semper nec. Nulla tempus quam tristique ipsum facilisis, nec gravida diam mattis.\n\nVivamus lorem ex, efficitur sed justo consectetur, luctus dignissim orci. Vestibulum rutrum pharetra elit, in sollicitudin tellus euismod eget. Nam nec aliquam metus. Suspendisse tristique ac massa ut fermentum. Cras id mi eu sapien rutrum ultricies. Phasellus id augue in elit hendrerit ornare ac in nisl. In viverra leo vel elit posuere, in tincidunt arcu sodales. Phasellus commodo justo ac egestas pharetra. Phasellus at ante ut dui egestas faucibus."
)

// CutSymbolDuplication

func testCutSymbolDuplication(beforeCut, afterCut []rune, t *testing.T) {
	if result := cutSymbolDuplication(' ', beforeCut); string(result) != string(afterCut) {
		t.Fatalf("expected: %s; fact: %s", string(afterCut), string(result))
	}
}

func benchmarkCutSymbolDuplication(sourceString []rune, b *testing.B) {
	for n := 0; n < b.N; n++ {
		cutSymbolDuplication(' ', sourceString)
	}
}

func TestCutSymbolDuplication1(t *testing.T) {
	testCutSymbolDuplication([]rune(test1BeforeCut), []rune(test1AfterCut), t)
}

func TestCutSymbolDuplication2(t *testing.T) {
	testCutSymbolDuplication([]rune(test2BeforeCut), []rune(test2AfterCut), t)
}

func TestCutSymbolDuplication3(t *testing.T) {
	testCutSymbolDuplication([]rune(test3BeforeCut), []rune(test3AfterCut), t)
}

func BenchmarkCutSymbolDuplication1(b *testing.B) {
	benchmarkCutSymbolDuplication([]rune(test1BeforeCut), b)
}

func BenchmarkCutSymbolDuplication2(b *testing.B) {
	benchmarkCutSymbolDuplication([]rune(test2BeforeCut), b)
}

func BenchmarkCutSymbolDuplication3(b *testing.B) {
	benchmarkCutSymbolDuplication([]rune(test3BeforeCut), b)
}