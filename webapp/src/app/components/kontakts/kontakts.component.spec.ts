import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { KontaktsComponent } from './kontakts.component';

describe('KontaktsComponent', () => {
  let component: KontaktsComponent;
  let fixture: ComponentFixture<KontaktsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KontaktsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KontaktsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
