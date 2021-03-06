function Chemin(p) {
	this.points = new Array();
	this.points.push(p);
	this.rect = new Rect(p.x, p.y, 1, 1);
	this.screenRect = new Rect();
}
Chemin.prototype = new ChCol();

Chemin.prototype.add = function(p) {
	var lastPoint = this.points[this.points.length-1];
	var dx = p.x-lastPoint.x;
	var dy = p.y-lastPoint.y;
	var adx = Math.abs(dx);
	var ady = Math.abs(dy);
	if (adx>ady) {
		this.points.push(new Point(dx>0 ? lastPoint.x+ady : lastPoint.x-ady, p.y));
	} else {
		this.points.push(new Point(p.x, dy>0 ? lastPoint.y+adx : lastPoint.y-adx));
	}
	this.points.push(p);
	this.rect.makeContain(p);
	return this;
}

Chemin.prototype.computeScreenRect = function(carte) {
	this.screenRect.x = carte.zoom*(carte.originX+this.rect.x);
	this.screenRect.y = carte.zoom*(carte.originY-this.rect.y);
	this.screenRect.w = carte.zoom*this.rect.w;
	this.screenRect.h = carte.zoom*this.rect.h;
}


Chemin.prototype.draw = function(carte) {
	carte.context.strokeStyle = '#e61803';//"#c99b1f";
	//~ this.screenRect.drawThin(carte.context);
	carte.context.lineWidth = Math.sqrt(carte.zoom);
	carte.context.lineCap = "round";
	carte.context.lineJoin = "round";
	carte.context.beginPath();
	var sp = new Point();
	carte.naturalToScreen(this.points[0], sp);
	carte.context.moveTo(sp.x, sp.y);
	for (var i=1; i<this.points.length; i++) {
		carte.naturalToScreen(this.points[i], sp);
		carte.context.lineTo(sp.x, sp.y);
	}
	carte.context.stroke();
}
